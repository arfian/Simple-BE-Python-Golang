import tornado.web
import os
import sys
import logging
import asyncio
from controllers import fetch

def run(port):
	application = tornado.web.Application([
		(r"/fetch/.*", fetch.FetchController),
	])
	if sys.platform == 'win32':
		asyncio.set_event_loop_policy(asyncio.WindowsSelectorEventLoopPolicy())
	application.listen(port)