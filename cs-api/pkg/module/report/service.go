package report

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"fmt"
	"time"
)

type service struct {
	repo iface.IRepository
}

func (s *service) ListDailyTagReport(ctx context.Context, startDate, endDate time.Time) ([]pkg.DailyTagReportColumn, map[string]map[string]int32, error) {
	tags, err := s.repo.GetAllTag(ctx)
	if err != nil {
		return nil, nil, err
	}

	var columns []pkg.DailyTagReportColumn
	for _, tag := range tags {
		columns = append(columns, pkg.DailyTagReportColumn{
			Label: tag.Name,
			Key:   fmt.Sprintf("tag_%d", tag.ID),
		})
	}

	// map[date]map[tag_x]count 每個日期會對應到一個 tag_{tag_id} 和 count 的 map
	resultMap := map[string]map[string]int32{}
	for d := startDate; d.After(endDate) == false; d = d.AddDate(0, 0, 1) {
		dateStr := d.Format("2006-01-02")
		if _, ok := resultMap[dateStr]; !ok {
			resultMap[dateStr] = make(map[string]int32)
			for _, tag := range tags {
				key := fmt.Sprintf("tag_%d", tag.ID)
				resultMap[dateStr][key] = 0
			}
		}
	}

	items, err := s.repo.ListReportDailyTag(ctx, model.ListReportDailyTagParams{
		Date:   startDate,
		Date_2: endDate,
	})
	if err != nil {
		return nil, nil, err
	}

	for _, item := range items {
		dateStr := item.Date.Format("2006-01-02")
		key := fmt.Sprintf("tag_%d", item.TagID)
		if v, ok := resultMap[dateStr]; ok {
			if _, ok2 := v[key]; ok2 {
				resultMap[dateStr][key] = item.Count
			}
		}
	}

	return columns, resultMap, nil
}

func (s *service) ListDailyStaffReport(ctx context.Context, startDate, endDate time.Time) ([]model.ListReportDailyStaffRow, error) {
	return s.repo.ListReportDailyStaff(ctx, model.ListReportDailyStaffParams{
		Date:   startDate,
		Date_2: endDate,
	})
}

func NewService(Repo iface.IRepository) iface.IReportService {
	return &service{
		repo: Repo,
	}
}
