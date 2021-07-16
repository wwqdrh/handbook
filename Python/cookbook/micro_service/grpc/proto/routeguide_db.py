# 用来处理中间数据源

import json
from micro_service.grpc.proto import routeguide_pb2
import os

JSON_PATH = os.path.join(os.path.dirname(__file__), "routeguide_db.json")


def read_routeguide_db():
    feature_list = []
    with open(JSON_PATH) as f:
        for item in json.load(f):
            feature = routeguide_pb2.Feature(
                name=item["name"],
                location=routeguide_pb2.Point(
                    latitude=item["location"]["latitude"],
                    longitude=item["location"]["longitude"],
                ),
            )
            feature_list.append(feature)
    return feature_list