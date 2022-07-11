local staffAssignRoom = function(id)
    local score = redis.call("ZSCORE", "zset:staff:dispatch", id)
    if tonumber(score) < 0 then
        return true
    end
    redis.call("ZADD", "zset:staff:dispatch", "XX", "INCR", 1, id)
    return true
end

return staffAssignRoom(KEYS[1])