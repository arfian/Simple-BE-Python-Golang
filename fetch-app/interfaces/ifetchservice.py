import abc

class IFetchService(metaclass=abc.ABCMeta):

    @abc.abstractmethod
    def changePrice(self):
        return