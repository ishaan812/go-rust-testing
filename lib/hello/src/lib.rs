extern crate zen_engine;
extern crate serde_json;
extern crate libc;
use std::collections::HashMap;
use std::ffi::{CStr, CString};
use serde_json::{json, Value};
use zen_engine::{DecisionEngine, DecisionGraphResponse};
use zen_engine::model::DecisionContent;
use libc::c_char;


#[derive(Debug)]
struct MyError {
    message: String,
}

#[no_mangle]
pub extern "C" fn evaluate(decision_content: *const libc::c_char, params: *const libc::c_char) -> *const c_char {
    let res = tokio::runtime::Runtime::new().unwrap().block_on(evaluate_decision(decision_content, params));
    match res {
        Ok(result) => {
            let result_str = serde_json::to_string(&result).unwrap();
            return CString::new(result_str).unwrap().into_raw()
        }
        Err(err) => {
            return CString::new(err.message).unwrap().into_raw()
        }
    };
}

async fn evaluate_decision(decision_content: *const c_char, params: *const c_char) -> Result<DecisionGraphResponse, MyError> {
    if decision_content.is_null() || params.is_null() {
        println!("Null pointer error");
        return Err(MyError {
            message: "Null pointer error".to_string(),
        });
    }
    let decision_content_str = unsafe { CStr::from_ptr(decision_content).to_string_lossy().into_owned() };
    let params_str = unsafe { CStr::from_ptr(params).to_string_lossy().into_owned() };
    let decision_content: DecisionContent = match serde_json::from_str(&decision_content_str) {
        Ok(content) => content,
        Err(err) => {
            println!("Error deserializing decision content: {}", err);
            return Err(MyError {
                message: "Error deserializing decision content".to_string(),
            });
        }
    };
    let map: HashMap<String, Value> = serde_json::from_str(&params_str).unwrap(); 
    let engine = DecisionEngine::default();
    let decision = engine.create_decision(decision_content.into());
    let result = decision.evaluate(&json!(map)).await;
    Ok(result.unwrap())
}