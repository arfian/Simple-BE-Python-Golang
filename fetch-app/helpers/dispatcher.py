import tornado.web
import config
import logging
import json
from datetime import datetime
import jwt

def delist_arguments(args):
    for arg, value in args.items():
        if len(value) == 1:
            args[arg] = value[0]
    return args

def checktoken(self):
    try:
        data = self.request.headers.get('Authorization')
        token = str.replace(str(data), 'Bearer ', '')
        return {'data': jwt.decode(token, "efishery123", algorithms=["HS256"]), 'code': 200}
    except Exception as exc:
        res = {'message': 'Token is missing', 'code': 401}
        return res

class MethodDispatcher(tornado.web.RequestHandler):
    def _dispatch(self):
        args = None
        # if self.request.arguments:
        #     args = delist_arguments(self.request.arguments)
        if self.request.uri.endswith('/'):
            func = getattr(self, 'index', None)
            if args:
                return func(**args)
            else:
                return func()
        path = self.request.uri.split('?')[0]
        method = path.split('/')[-1]
        if not method.startswith('_'):
            func = getattr(self, method, None)
            if func:
                if args:
                    return func(**args)
                else:
                    return func()
            else:
                raise tornado.web.HTTPError(404)
        else:
            raise tornado.web.HTTPError(404)

    def get(self):
        return self._dispatch()

    def post(self):
        return self._dispatch()

    # Jika tidak dibutuhkan comment saja ini
    def prepare(self):
        return super(MethodDispatcher, self).prepare()

    def on_finish(self):
        return super(MethodDispatcher, self).on_finish()