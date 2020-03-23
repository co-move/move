use std::ffi::{
    CStr,
    CString,
};
use std::os::raw::c_char;
mod vm;

#[no_mangle]
pub extern "C" fn hello(name: *const c_char) -> *const c_char {
    unsafe {
        return CStr::from_ptr(name).as_ptr();
    }
}

#[no_mangle]
pub unsafe extern "C" fn run_script(scriptc: *mut c_char) -> *const c_char {
    let script = CString::from_raw(scriptc );
    println!("script received in rust: {:?}", script);
    let output = vm::run_script(hex::decode(script.into_string().unwrap()).unwrap().to_vec());
    println!("{:?}",output);
    let cstr = CString::new(output).expect("CString Convert failed");
    cstr.into_raw()
}
