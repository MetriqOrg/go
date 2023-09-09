# dump-ledger-state

This tool dumps the state from history archive buckets to 4 separate files:
* accounts.csv
* accountdata.csv
* offers.csv
* trustlines.csv
* claimablebalances.csv

It's primary use is to test `SingleLedgerStateReader`. To run the test (`run_test.sh`) it:
1. Runs `dump-ledger-state`.
2. Syncs gramr to the same checkpoint: `gramr catchup [ledger]/1`.
3. Dumps gramr DB by using `dump_core_db.sh` script.
4. Diffs results by using `diff_test.sh` script.
