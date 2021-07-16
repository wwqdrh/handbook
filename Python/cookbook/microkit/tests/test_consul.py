from microkit import consul
import uuid


def clean_consul():
    consul_client = consul.Consul("172.27.104.48:8500")
    server_ids = []
    for server_id in server_ids:
        consul_client.deregister(server_id)


def test_consul_register():
    from sanic import Sanic, response
    import atexit

    app = Sanic(__name__)

    @app.route("/health")
    async def health(request):
        return response.json({"status": "ok"})

    consul_client = consul.Consul("172.27.104.48:8500")
    service_id = "test_service_{}".format(str(uuid.uuid4()).replace("-", ""))
    consul_client.register(
        service_id,
        "test_service",
        "172.27.104.48",
        8082,
        "primary,test",
    )

    atexit.register(
        lambda client, server_id: client.deregister(server_id),
        consul_client,
        service_id,
    )

    app.run("0.0.0.0", 8082)


test_consul_register()
# clean_consul()