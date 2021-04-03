from interfaces import ifetchservice
import logging
from cachetools import cached, TTLCache
import config

class FetchService(ifetchservice.IFetchService):

    def __init__(self, repo):
        self.repo = repo

    def changePrice(self):
        res = self.repo.fetchResources()
        conversion = 0

        try:
            conversion = config.app['cache']['convert']
        except Exception as exc:
            logging.info('request conversi')
            resConvert = self.repo.fetchConversion()
            config.app['cache']['convert'] = resConvert["USD_IDR"]
        
        logging.info(resConvert)
        for i in res:
            if i["price"] != None:
                i["price_usd"] = float(i["price"]) / resConvert["USD_IDR"]
            else:
                i["price_usd"] = 0
        
        return res