# from redis.client import Redis,StrictRedis
# from redis.connection import (
#     BlockingConnectionPool,
#     ConnectionPool,
#     Connection,
#     SSLConnection,
#     UnixDomainSocketConnection
# )
# from redis.utils import from_url
# from redis.exceptions import (
#     AuthenticationError,
#     BusyLoadingError,
#     ConnectionError,
#     DataError,
#     InvalidResponse,
#     PubSubError,
#     ReadOnlyError,
#     RedisError,
#     ResponseError,
#     TimeoutError,
#     WatchError
# )

from redis import *

r = StrictRedis(host='127.0.0.1',port='6379')
# 写
pipe = r.pipeline()
pipe.set('py1','hello1')
pipe.set('py2','world')
pipe.execute()

temp = r.get('py1')
print(temp)