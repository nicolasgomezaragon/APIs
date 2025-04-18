class MetaData:
    def __init__(self, information, symbol, last_refreshed, output_size, time_zone):
        self.information = information
        self.symbol = symbol
        self.last_refreshed = last_refreshed
        self.output_size = output_size
        self.time_zone = time_zone

class DailyData:
    def __init__(self, open, high, low, close, volume):
        self.open = open
        self.high = high
        self.low = low
        self.close = close
        self.volume = volume

    def to_dict(self):
        return {
            "open": float(self.open),
            "high": float(self.high),
            "low": float(self.low),
            "close": float(self.close),
            "volume": int(self.volume),
        }

class TimeSeries:
    def __init__(self, meta_data, time_series):
        self.meta_data = MetaData(**meta_data)
        self.time_series = {
            date: DailyData(**data)
            for date, data in time_series.items()
        }
