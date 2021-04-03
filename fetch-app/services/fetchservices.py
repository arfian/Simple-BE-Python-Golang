from interfaces import ifetchservice

class FetchService(ifetchservice.IFetchService):

    def __init__(self, repo):
        self.repo = repo

    def changePrice(self):
        res = self.repo.fetchResources()
        return res