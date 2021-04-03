import abc

class IFetchRepo(metaclass=abc.ABCMeta):

    @abc.abstractmethod
    def fetchResources(self):
        return