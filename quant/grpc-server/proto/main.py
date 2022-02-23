import datetime
import json
from concurrent import futures

import grpc
from google.protobuf.timestamp_pb2 import Timestamp
from google.protobuf.json_format import MessageToJson

import quant_pb2
import quant_pb2_grpc
from calculate import *
from datetime import datetime, timezone


def dic_to_list(obj):
    res = [obj["min"], obj["max"]]
    return res


class Quant(quant_pb2_grpc.QuantServicer):
    def Request(self, request, context):
        print('======= Got request =======')
        json_obj = MessageToJson(request, including_default_value_fields=True)
        req = json.loads(json_obj)

        date_format = "%Y-%m-%dT%H:%M:%SZ"
        if req["start_date"] == '0001-01-01T00:00:00Z':
            req["start_date"] = None
        else:
            req["start_date"] = datetime.strptime(req["start_date"], date_format).replace(tzinfo=timezone.utc)

        if req["end_date"] == '0001-01-01T00:00:00Z':
            req["end_date"] = None
        else:
            req["end_date"] = datetime.strptime(req["end_date"], date_format).replace(tzinfo=timezone.utc)

        # net_revenue
        # net_profit
        # market_cap
        req["net_revenue"] = {'min': int(req["net_revenue"]["min"]), 'max': int(req["net_revenue"]["max"])}
        req["net_profit"] = {'min': int(req["net_profit"]["min"]), 'max': int(req["net_profit"]["max"])}
        req["market_cap"] = {'min': int(req["market_cap"]["min"]), 'max': int(req["market_cap"]["max"])}

        print("request: ", req)
        try:
            start = time.time()
            find_code = FindCode()
            code_list = find_code.apply_conditions(
                start_date=datetime(2016, 12, 30, 0, 0, 0, tzinfo=timezone.utc), end_date=None,
                term=12, market=None, main_sector=["IT"], net_rev=[10000, 1000000000],
                net_rev_r=[None, None], net_prf=[None, None], net_prf_r=[None, None], de_r=[None, None],
                per=[0, 10], psr=[None, None], pbr=[0, 10], pcr=[None, None], op_act=[0, 1000000], iv_act=[-1000000, 0],
                fn_act=[None, None], dv_yld=[None, None], dv_pay_r=[None, None], roa=[None, None], roe=[None, None],
                marcap=[None, None]
                # start_date=req["start_date"],
                # end_date=req["end_date"],
                # term=12,
                # market=None,
                # main_sector=req["main_sector"],
                # net_rev=dic_to_list(req["net_revenue"]),
                # # net_rev_r=dic_to_list(req["net_revenue_rate"]),
                # net_rev_r=[None, None],
                # net_prf=dic_to_list(req["net_profit"]),
                # # net_prf_r=dic_to_list(req["net_profit_rate"]),
                # net_prf_r=[None, None],
                # de_r=dic_to_list(req["de_ratio"]),
                # per=dic_to_list(req["per"]),
                # # psr=dic_to_list(req["psr"]),
                # psr=[None, None],
                # pbr=dic_to_list(req["pbr"]),
                # # pcr=dic_to_list(req["pcr"]),
                # pcr=[None, None],
                # op_act=dic_to_list(req["activities"]["operating"]),
                # iv_act=dic_to_list(req["activities"]["investing"]),
                # fn_act=dic_to_list(req["activities"]["financing"]),
                # dv_yld=dic_to_list(req["dividend_yield"]),
                # dv_pay_r=dic_to_list(req["dividend_payout_ratio"]),
                # roa=dic_to_list(req["roa"]),
                # roe=dic_to_list(req["roe"]),
                # marcap=dic_to_list(req["market_cap"])
            )
            calculate = Calculate()
            return_dict = calculate.calculate_profit(code_list)
            print(return_dict)
            delta_t = time.time() - start
            print("total process : ", delta_t, "s")

            dt = return_dict["chart"]["start_date"]
            t = dt.timestamp()
            seconds = int(t)
            nanos = int(t % 1 * 1e9)
            proto_timestamp = Timestamp(seconds=seconds, nanos=nanos)
            return_dict["chart"]["start_date"] = proto_timestamp

            return quant_pb2.QuantResult(
                cumulative_return=return_dict["cumulative_return"],
                annual_average_return=return_dict["annual_average_return"],
                winning_percentage=return_dict["winning_percentage"],
                max_loss_rate=return_dict["max_loss_rate"],
                holdings_count=return_dict["holdings_count"],
                chart_data=return_dict["chart"],
            )

        except Exception as e:
            print(e)
            return


def serve():
    print('gRPC server started. Listening to 9000...')
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    quant_pb2_grpc.add_QuantServicer_to_server(Quant(), server)

    server.add_insecure_port('[::]:9000')

    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    print('=' * 20)
    serve()
