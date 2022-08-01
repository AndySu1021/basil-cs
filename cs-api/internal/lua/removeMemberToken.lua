local removeMemberToken = function(name)
    local ns = "token:member:"
    local token = redis.call("GET", ns .. name)
    if token ~= false then
        redis.call("DEL", ns .. token)
        redis.call("DEL", ns .. name)
    end
    return true
end

return removeMemberToken(KEYS[1])