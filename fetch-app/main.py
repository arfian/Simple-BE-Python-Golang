import tornado.ioloop
import signal
import logging
from tornado.options import options
from routers import router
import config
from helpers import dispatcher

is_closing = False

def signal_handler(signum, frame):
	global is_closing
	logging.info('exiting...')
	is_closing = True

def try_exit(): 
	global is_closing
	if is_closing:
		tornado.ioloop.IOLoop.instance().stop()
		logging.info('exit success')

if __name__ == "__main__":
	tornado.options.parse_command_line()
	signal.signal(signal.SIGINT, signal_handler)
	router.run(config.app['httpport'])
	logging.info('start...' + config.app['httpport'])
	tornado.ioloop.PeriodicCallback(try_exit, 100).start() 
	tornado.ioloop.IOLoop.instance().start()