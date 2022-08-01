local staffCloseRoom = function(staffId, memberName)
    local ns = "token:member:"
    local token = redis.call("GET", ns .. memberName)
    if token ~= false then
        redis.call("DEL", ns .. token)
        redis.call("DEL", ns .. memberName)
    end
    redis.call("ZADD", "zset:staff:dispatch", "XX", "INCR", -1, staffId)
    return true
end

return staffCloseRoom(KEYS[1], KEYS[2])