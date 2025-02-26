# Blockchain Consensus Mechanism

Blockchains are maintained through a "consensus" mechanism, where all nodes must agree upon the state of the blockchain at all times. This ensures that all participants in the network maintain an identical view of the blockchain state.

## Consensus Process

When a transaction is accepted, each node calculates a new state for the blockchain and seeks "consensus" with other nodes. The consensus process involves the following steps:

1. **Transaction Validation:** Nodes validate incoming transactions based on predefined rules.
2. **State Calculation:** Each node computes a new state after applying the validated transactions.
3. **Consensus:** Nodes compare their computed state and seek agreement on the blockchain's state.

## Consensus-Breaking Changes

In some scenarios, changes to the blockchain can break consensus. A **consensus-breaking change** occurs when nodes calculate different states from the same set of transactions.

### Adding Fields to Resources

 My change is that a new field is added to the `Ticket` resource, nodes running older versions of the blockchain software will not be aware of this new field. As a result, they will not hash this new field, leading to a different state calculation compared to nodes that have visibility of the new field.

This inconsistency can cause the blockchain to fork, where nodes diverge into different versions of the chain with conflicting states.