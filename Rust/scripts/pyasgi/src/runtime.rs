use once_cell::sync::Lazy;
use tokio::runtime::Runtime;

use std::time::Duration;

use pyo3::prelude::*;

pub static RUNTIME: Lazy<Runtime> = Lazy::new(|| Runtime::new().unwrap());

const PYTHON_CODE: &'static str = r#"
 import asyncio

 async def py_sleep(duration):
     await asyncio.sleep(duration)
 "#;

async fn py_sleep(seconds: f32) -> PyResult<()> {
    let test_mod = Python::with_gil(|py| -> PyResult<PyObject> {
        Ok(
            PyModule::from_code(py, PYTHON_CODE, "test_into_future/test_mod.py", "test_mod")?
                .into(),
        )
    })?;

    Python::with_gil(|py| {
        pyo3_asyncio::into_future(
            test_mod
                .call_method1(py, "py_sleep", (seconds.into_py(py),))?
                .as_ref(py),
        )
    })?
    .await?;
    Ok(())
}
