# Comove Demo

This is a demo that runs a simple Libra Move script on Cosmos.

The [sample script](./scripts/test.mv) is compiled from this source file.
```
main() {
    assert(0u64 + 0u64 == 0u64, 1000);
    assert(0u64 + 1u64 == 1u64, 1001);
    assert(1u64 + 1u64 == 2u64, 1002);
    return;
}
```

# Getting start

```
movecli tx move run-script {{path_to_script}}/test.mv 123 --from ping --chain-id test -y
```
