#[allow(unused_imports)]
use language_e2e_tests::data_store::FakeDataStore;
use libra_types::{
    //account_address::AccountAddress,
    account_config,
};
use move_vm_runtime::MoveVM;
use move_vm_state::execution_context::SystemExecutionContext;
use stdlib::{stdlib_modules, StdLibOptions};
use vm::{
    errors::VMResult,
    //access::ModuleAccess,
    gas_schedule::{
        //AbstractMemorySize,
        CostTable,
        GasAlgebra,
        //GasCarrier,
        GasUnits,
    },
    transaction_metadata::TransactionMetadata,
    // file_format::{
    //     //CompiledModule, 
    //     CompiledScript},
};

#[allow(dead_code)]
pub fn run_script(raw: Vec<u8>) -> String {
    let address = account_config::association_address(); //AccountAddress::default();

    // Execute script.
    // create a Move VM and populate it with generated modules
    let move_vm = MoveVM::new();
    let data_cache = FakeDataStore::default();
    let mut ctx = SystemExecutionContext::new(&data_cache, GasUnits::new(0));
    let gas_schedule = CostTable::zero();

    // load std modules
    let mut txn_stdlib = TransactionMetadata::default();
    txn_stdlib.sender = account_config::CORE_CODE_ADDRESS;
    let std = stdlib_modules(StdLibOptions::Staged).iter();
    for x in std {
        let mut bytes: Vec<u8> = vec![];
        x.serialize(&mut bytes)
            .expect("Std Modules serialize failed.");
        move_vm
            .publish_module(bytes, &mut ctx, &txn_stdlib)
            .expect("Publish failed");
    }

    let mut txn_data = TransactionMetadata::default();
    txn_data.sender = address;

    let result = move_vm.execute_script(raw, &gas_schedule, &mut ctx, &txn_data, vec![]);

    result_to_json(result)
}

fn result_to_json(result: VMResult<()>) -> String {
    match result {
        Ok(_) => "{\"code\":0, \"msg\":\"\"}".to_string(),
        Err(e) => format!("{{\"code\":101, \"msg\":\"{:?}\"}}", e).to_string()
    }
}