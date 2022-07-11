local staffTransferRoom = function(staffId, toStaffId)
    redis.call("ZADD", "zset:staff:dispatch", "XX", "INCR", -1, staffId)
    local score = redis.call("ZSCORE", "zset:staff:dispatch", toStaffId)
    if tonumber(score) < 0 then
        return true
    end
    redis.call("ZADD", "zset:staff:dispatch", "XX", "INCR", 1, toStaffId)
    return true
end

return staffTransferRoom(KEYS[1], KEYS[2])