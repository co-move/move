# Comove Demo

This is demo that runs a simple Move script on Cosmos.

The sample script is compiled from this source file.
```
main() {
    assert(0u64 + 0u64 == 0u64, 1000);
    assert(0u64 + 1u64 == 1u64, 1001);
    assert(1u64 + 1u64 == 2u64, 1002);
    return;
}
```