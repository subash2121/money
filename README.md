# Money

A package to process indian currency in rupees or paise to add values of money and check for equality.

## Problem Statement

> As a fan of wealth,
> I want to model money which can be expressed in Rupees, paise or a combination thereof,
> So that I can add values of money.

## Prerequisites

- Go (1.18.2)

## Requirements

- https://github.com/stretchr/testify

## Test Instructions

Execute the below command 

    $ go test ./...

## Build Instructions

Execute the below command to build the package

    $ go build ./...

## Library Usage

To create a new money

    <variableOne> = NewMoney()
    <variableTwo> = NewMoney()

To add paise to existing money

    <variableOne>.AddPaise(<amount>)

To subtract paise to existing money

    <variableOne>.SubtractPaise(<amount>)

To add rupee to existing money

    <variableOne>.AddRupee(<amount>)

To subtract rupee to existing money

    <variableOne>.SubtractRupee(<amount>)

To add two money

    <variableOne>.AddMoney(<variableTwo>)

To subtract two money

    <variableOne>.SubtractMoney(<variableTwo>)

To check for equality of two money (returns boolean value)

    <variableOne>.EqualsTo(<variableTwo>)