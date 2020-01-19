CREATE DATABASE IF NOT EXISTS accounts CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE accounts;

CREATE TABLE accounts (
    id              INT(10) NOT NULL AUTO_INCREMENT,   --
    balance         BIGINT NOT NULL,                   -- int64 -9223372036854775807 - 9223372036854775807
    version         TINYINT UNSIGNED NOT NULL,         -- uin8 0-255 optimistic concurrency control (OCC)
    user_id         INT(10) NOT NULL,                  --
    type            CHAR(4) NOT NULL,                  -- CASH, COIN
    is_withdrawable BOOLEAN NOT NULL,                  -- is possible to withdraw to own bank account
    PRIMARY KEY (id)
);

/*
    ATM: Deposit or withdraw funds using an ATM.
    Charge: Record a purchase on a credit card or withdraw funds using a debit card.
    Check: Withdraw funds by writing a paper check. Choosing this type will automatically insert a number in the '#' field (the next number in sequence from the last check recorded).
    Deposit: Add funds to an account by any method.
    Online: Withdraw funds through a web-based store or online banking service.
    POS: Withdraw funds through a point-of-sale transaction (typically a cash or debit card purchase).
    Transfer: Move funds from one account to another (for more information, see Account Transfers).
    Withdrawal: Deduct funds from an account by any method.
    Withdraw:
 */

CREATE TABLE txs (
    id           INT(10) NOT NULL AUTO_INCREMENT,       --
    tx_type      CHAR(10) NOT NULL,                     -- withdraw, deposit, charge
    created_at   DATETIME(3) NOT NULL,                  --
    amount       BIGINT NOT NULL,                       -- int16 -9223372036854775807 - 9223372036854775807
    accounts_id  INT(10) NOT NULL,                      --
    FOREIGN KEY (accounts_id) REFERENCES accounts (id), --
    PRIMARY KEY (id)
);