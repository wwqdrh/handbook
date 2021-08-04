# python setup.py build_ext --inplace

from distutils.core import setup
from Cython.Build import cythonize

setup(
    name="app",
    ext_modules=cythonize(
        "func.py",
        compiler_directives={"language_level": "3"},
    ),
)
