local removeStaffToken = function(id, username)
    local ns = "token:staff:"
    local token = redis.call("GET", ns .. username)
    if token ~= false then
        redis.call("DEL", ns .. token)
        redis.call("DEL", ns .. username)
    end
    redis.call("ZREM", "zset:staff:dispatch", id)
    return true
end

return removeStaffToken(KEYS[1], KEYS[2])