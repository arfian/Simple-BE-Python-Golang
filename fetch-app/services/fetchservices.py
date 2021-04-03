from interfaces import ifetchservice
import logging
from cachetools import cached, TTLCache
import config
from itertools import groupby
import json
import numpy as np
import pandas as pd

class FetchService(ifetchservice.IFetchService):

    def __init__(self, repo):
        self.repo = repo

    def changePrice(self):
        res = self.repo.fetchResources()
        conversion = 0

        try:
            conversion = config.app['cache']['convert']
            logging.info("no cache")
        except Exception as exc:
            resConvert = self.repo.fetchConversion()
            config.app['cache']['convert'] = resConvert["USD_IDR"]
            conversion = resConvert["USD_IDR"]
            logging.info("cache")
        
        for i in res:
            if i["price"] != None:
                i["price_usd"] = float(i["price"]) / conversion
            else:
                i["price_usd"] = 0
        
        return res

    def calAggregate(self):
        res = self.repo.fetchResources()
        res.sort(key=lambda x: 'kosong' if x['area_provinsi'] is None else x['area_provinsi'])
        groups = groupby(res, lambda x: 'kosong' if x['area_provinsi'] is None else x['area_provinsi'])
        areadata = []
        for area, group in groups:
            if area != 'kosong':
                areadata.append({'area_provinsi': area, 'count': len(list(group))})
        arealengths = [x['count'] for x in areadata]

        d_list = [item for item in res if item['tgl_parsed'] != None]
        df = pd.DataFrame(d_list)
        df['timestamp'] = pd.to_datetime(df['timestamp'], unit='ms')

        df.sort_values(by='timestamp')
        gr = groupby(df['timestamp'], lambda x: x.week)

        weekdata = []
        for name, group in gr:
            weekdata.append({'week': name, 'count': len(list(group))})
        weeklengths = [x['count'] for x in weekdata]

        resdata = {
            'aggregate_area': {
                'max': max(areadata, key=lambda x: x['count']),
                'min': min(areadata, key=lambda x: x['count']),
                'average': sum(x['count'] for x in areadata) / len(areadata),
                'median': np.median(arealengths),
            },
            'aggregate_week': {
                'max': max(weekdata, key=lambda x: x['count']),
                'min': min(weekdata, key=lambda x: x['count']),
                'average': sum(x['count'] for x in weekdata) / len(weekdata),
                'median': np.median(weeklengths),
            },
            'data_area_provinsi': areadata,
            'data_week': weekdata
        }

        return resdata