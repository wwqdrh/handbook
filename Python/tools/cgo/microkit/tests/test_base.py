import microkit


def test_package_info():
    assert microkit.PACKAGE_NAME == "microkit"
    # microkit.PACKAGE_PATH # /.../microkit


test_package_info()