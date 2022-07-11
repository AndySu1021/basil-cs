import httpRequest from "@/utils/http";

export const apiGetStaffList = async (query) => {
  return httpRequest({
    url: '/staffs?' + query,
    method: 'GET',
  })
}

export const apiCreateStaff = async (data) => {
  return httpRequest({
    url: '/staff',
    method: 'POST',
    data
  })
}

export const apiGetStaffDetail = async (id) => {
  return httpRequest({
    url: '/staff/' + id,
    method: 'GET',
  })
}

export const apiUpdateStaff = async (id, data) => {
  return httpRequest({
    url: '/staff/' + id,
    method: 'PUT',
    data
  })
}

export const apiDeleteStaff = async (id) => {
  return httpRequest({
    url: '/staff/' + id,
    method: 'DELETE',
  })
}

export const apiUpdateStaffServingStatus = async (data) => {
  return httpRequest({
    url: '/staff/serving-status',
    method: 'PATCH',
    data
  })
}

export const apiUpdateStaffAvatar = async (data) => {
  return httpRequest({
    url: '/staff/avatar',
    method: 'PATCH',
    data
  })
}

export const apiGetAvailableStaff = async () => {
  return httpRequest({
    url: '/available-staffs',
    method: 'GET',
  })
}

export const apiGetAllStaff = async () => {
  return httpRequest({
    url: '/staffs/all',
    method: 'GET',
  })
}
