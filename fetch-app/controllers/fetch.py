import tornado.web
import json
from helpers import dispatcher
from services import fetchservices
from repositories import fetchrepository
import logging

class FetchController(dispatcher.MethodDispatcher):
	
	def getprice(self):
		checkdata = dispatcher.checktoken(self)
		if checkdata["code"] == 200:
			repo = fetchrepository.FetchRepo()
			services = fetchservices.FetchService(repo)
			resData = services.changePrice()
			self.write({'code': 200, 'message': 'Berhasil', 'data': resData})
		else:
			self.write(checkdata)

	def getaggregate(self):
		checkdata = dispatcher.checktoken(self)
		if checkdata["code"] == 200:
			repo = fetchrepository.FetchRepo()
			services = fetchservices.FetchService(repo)
			resData = services.calAggregate()
			self.write({'code': 200, 'message': 'Berhasil', 'data': resData})
		else:
			self.write(checkdata)
