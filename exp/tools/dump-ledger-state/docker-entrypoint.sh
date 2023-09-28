#! /bin/bash
set -e

/etc/init.d/postgresql start

while ! psql -U circleci -d core -h localhost -p 5432 -c 'select 1' >/dev/null 2>&1; do
    echo "Waiting for postgres to be available..."
    sleep 1
done

echo "using version $(gravity version)"

if [ -z ${TESTNET+x} ]; then
    gravity --conf ./gravity.cfg new-db
else
    gravity --conf ./gravity-testnet.cfg new-db
fi

if [ -z ${LATEST_LEDGER+x} ]; then
    # Get latest ledger
    echo "Getting latest checkpoint ledger..."
    if [ -z ${TESTNET+x} ]; then
        export LATEST_LEDGER=`curl -s http://gravity8history.lantah.network/.well-known/lantah-history.json | jq -r '.currentLedger'`
    else
        export LATEST_LEDGER=`curl -s http://testgravity1history.lantah.network/.well-known/lantah-history.json | jq -r '.currentLedger'`
    fi
fi

if [[ -z "${LATEST_LEDGER}" ]]; then
  echo "could not obtain latest ledger"
  exit 1
fi

echo "Latest ledger: $LATEST_LEDGER"

if ! ./run_test.sh; then
    echo "ingestion dump (git commit \`$GITCOMMIT\`) of ledger \`$LATEST_LEDGER\` does not match gravity db."
    exit 1
fi