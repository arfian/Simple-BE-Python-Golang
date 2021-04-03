import abc

class IFetchRepo(metaclass=abc.ABCMeta):

    @abc.abstractmethod
    def fetchResources(self):
        return

    @abc.abstractmethod
    def fetchConversion(self):
        return