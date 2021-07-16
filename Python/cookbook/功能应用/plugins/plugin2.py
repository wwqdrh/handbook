from . import register_plugin

@register_plugin
def hello_2():
    print("Hello from Plugin 2")

@register_plugin
def goodbye():
    print("Plugin 2 says goodbye")