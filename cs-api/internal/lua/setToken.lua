local setToken = function(type, name, token, value, expire)
    local ns = "token:" .. type .. ":"
    redis.call("SET", ns .. name, token, "EX", expire)
    redis.call("SET", ns .. token, value, "EX", expire)
    return true
end

return setToken(KEYS[1], KEYS[2], KEYS[3], ARGV[1], ARGV[2])