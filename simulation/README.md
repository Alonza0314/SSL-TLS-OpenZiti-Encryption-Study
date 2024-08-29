# Ziti's Model Simulation

## Sequence Diagram

```mermaid
sequenceDiagram
    participant c as Client
    participant s as Server

    c -->> c: 1. KeyPairC, txHeaderC
    c ->> s: 2. clientHello
    s -->> s: 3. KeyPairS, rxS, txS, txHeaderS, senderS, receiverS
    s ->> c: 4. serverHello
    c -->> c: 5. rxC, txC, senderC, receiverC
    Note over c, s: Crypto Communication
```

## Crypto Model - Base Curve25519

```mermaid
graph TB
    subgraph Compute
        direction LR

        subgraph Server
            direction TB
            SA[SessionKey - skS] -->|+pkC| SC[SharePoint - q]
            SB[PublicKey - pkS] --> SD["Key = hash(q + pkC + pkS)"]
            SC --> SD
            SD -->|"Key[32:64]"| SE[rxS]
            SD -->|"Key[0:32]"| SF[txS]
        end

        subgraph Client
            direction TB
            CA[SessionKey - skC] -->|+pkC| CC[SharePoint - q]
            CB[PublicKey - pkC] --> CD["Key = hash(q + pkC + pkS)"]
            CC --> CD
            CD -->|"Key[0:32]"| CE[rxC]
            CD -->|"Key[32:64]"| CF[txC]
        end

        Server -->|pkS, txHeaderS| Client
        Client -->|pkC, txHeaderC| Server
    end

    subgraph Communication
        direction LR
        
        SG --- CH
        CG --- SH
    end

    Server --> |"(rxS)"| Communication
    Server --> |"(txS)"| Communication
    Client --> |"(rxC)"| Communication
    Client --> |"(txC)"| Communication
```
