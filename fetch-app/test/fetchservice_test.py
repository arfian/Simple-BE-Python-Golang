from services import fetchservices
from repositories import fetchrepository
import unittest
import config

class TestMethods(unittest.TestCase):

    def testChangePrice(self):
        repo = fetchrepository.FetchRepo()
        services = fetchservices.FetchService(repo)
        resData = services.changePrice()
        self.assertGreater(len(resData), 0)
    
    def testCalAggregate(self):
        repo = fetchrepository.FetchRepo()
        services = fetchservices.FetchService(repo)
        resData = services.calAggregate()
        self.assertGreater(len(resData['data_week']), 0)

if __name__ == '__main__':
    unittest.main()
