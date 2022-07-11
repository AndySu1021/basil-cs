local staffRegister = function(id)
    local score = redis.call("ZSCORE", "zset:staff:dispatch", id)
    if score == false then
        redis.call("ZADD", "zset:staff:dispatch", -1, id)
    end
    return true
end

return staffRegister(KEYS[1])