# x/mint Specification

## Overview

The `x/mint` module is responsible for minting new tokens and distributing them as rewards in the network. This module defines the minting schedule and manages the total supply of minted tokens. It ensures a smooth increase in the token supply according to predefined rules.

## State

### Minter

The `Minter` keeps track of the current annual inflation rate and the inflation rate applied during each block:

- `inflation`: The current annual rate of inflation.
- `annual_provisions`: The amount of tokens to be minted in the current year.

### Params

The `Params` define the minting parameters which can be modified by governance:

- `mint_denom`: The denomination of the minted tokens.
- `inflation_rate_change`: The annual increase or decrease of the inflation rate.
- `inflation_max`: The maximum inflation rate.
- `inflation_min`: The minimum inflation rate.
- `goal_bonded`: The desired ratio of bonded tokens to total supply.
- `blocks_per_year`: The expected number of blocks per year.

## Messages

### MsgUpdateParams

Governance can update the minting parameters using `MsgUpdateParams`.

## Events

### Mint

Emitted when new tokens are minted:

- `amount`: The amount of newly minted tokens.
- `annual_provision`: The annual provision applied.

## Queries

### QueryInflation

Returns the current inflation rate.

### QueryAnnualProvisions

Returns the current annual provisions.

## Genesis

The `GenesisState` defines the initial state of the `x/mint` module:

- `minter`: Initial minter state.
- `params`: Initial parameters.

## Parameters

Various parameters within the `x/mint` module can be modified via governance proposals. These changes are handled in `MsgUpdateParams`.

## BeginBlock

During each `BeginBlock`, the module checks the current block height and calculates the amount of newly minted tokens according to the inflation rate and updates the supply.

## EndBlock

At the end of each block, the new state is committed.

## Usage

To interact with the module, the following commands and queries are available:

### CLI Commands

- `mint query inflation`: Query the current inflation rate.
- `mint query annual-provisions`: Query the current annual provisions.
- `mint tx update-params`: Update the minting parameters via a governance proposal.

### REST Endpoints

- `GET /mint/inflation`: Query the current inflation rate.
- `GET /mint/annual_provisions`: Query the current annual provisions.