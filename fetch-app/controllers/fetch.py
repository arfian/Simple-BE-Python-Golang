import tornado.web
import json
from helpers import dispatcher
from services import fetchservices
from repositories import fetchrepository

class FetchController(dispatcher.MethodDispatcher):
	
	def getprice(self):
		repo = fetchrepository.FetchRepo()
		services = fetchservices.FetchService(repo)
		resData = services.changePrice()
		self.write({'code': 200, 'message': 'Berhasil', 'data': resData})
