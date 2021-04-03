from cachetools import cached, TTLCache

app = dict (
	httpport = '2002',
	apikey = 'cc281165ebd600247005',
	cache = TTLCache(maxsize=10, ttl=6000)
)