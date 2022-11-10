package main

import (
	"errors"
	"github.com/magiconair/properties"
	"regexp"
	"strings"
)

// Transition
/*
 *  Let δ(p,a) = q.
 *  p = CurrentState
 *  a = NextTerminal
 *  q = TargetState
 */
type Transition struct {
	CurrentState string
	NextTerminal string
	TargetState  string
}

func (tr Transition) String() string {

	return "δ(" + tr.CurrentState + ", " + tr.NextTerminal + ") = " + tr.TargetState
}

// FiniteAutomaton
/*
 *  Let M = (Q,Σ,δ,p,F).
 *  Q = States
 *  Σ = Alphabet
 *  δ = Transitions
 *  p = InitialState
 *  F = FinalStates
 */
type FiniteAutomaton struct {
	States       []string
	Alphabet     []string
	Transitions  []Transition
	InitialState string
	FinalStates  []string
}

func (fa *FiniteAutomaton) ReadFA(filepath string) error {

	p, err := properties.LoadFile(filepath, properties.UTF8)

	if err != nil {
		return err
	}

	keys := p.Keys()

	for _, key := range keys {

		switch key {

		case "states":
			fa.States = strings.Split(p.MustGetString(key), ",")
			break

		case "alphabet":
			fa.Alphabet = strings.Split(p.MustGetString(key), ",")
			break

		case "transitions":

			rawTransitions := strings.Split(p.MustGetString(key), ",")

			err = fa.parseTransitions(rawTransitions)

			if err != nil {
				fa.States = make([]string, 0)
				fa.Alphabet = make([]string, 0)
				fa.Transitions = make([]Transition, 0)
				fa.InitialState = ""
				fa.FinalStates = make([]string, 0)

				return err
			}

			break

		case "initialState":
			fa.InitialState = p.MustGetString(key)
			break

		case "finalStates":
			fa.FinalStates = strings.Split(p.MustGetString(key), ",")
			break

		default:

			fa.States = make([]string, 0)
			fa.Alphabet = make([]string, 0)
			fa.Transitions = make([]Transition, 0)
			fa.InitialState = ""
			fa.FinalStates = make([]string, 0)

			return errors.New("invalid finite FA parameter name: " + key)
		}
	}

	err = fa.validateTransitions()

	if err != nil {
		fa.States = make([]string, 0)
		fa.Alphabet = make([]string, 0)
		fa.Transitions = make([]Transition, 0)
		fa.InitialState = ""
		fa.FinalStates = make([]string, 0)

		return err
	}

	return nil
}

func (fa *FiniteAutomaton) Accepts(sequence string) bool {

	if !fa.validateSequence(sequence) {
		return false
	}

	var currentState = fa.InitialState

	for _, term := range sequence {

		validTransition := false

		for _, tr := range fa.Transitions {

			if currentState == tr.CurrentState && string(term) == tr.NextTerminal {

				currentState = tr.TargetState
				validTransition = true
				break
			}
		}

		if !validTransition {
			return false
		}
	}

	validFinalState := false

	for _, st := range fa.FinalStates {

		if currentState == st {
			validFinalState = true
			break
		}
	}

	return validFinalState
}

func (fa *FiniteAutomaton) parseTransitions(rawTransitions []string) error {

	var splitTerms []string
	var newTransition Transition

	for _, tr := range rawTransitions {

		// Checking if a transition read from a file is surrounded by a pair of parentheses, and between t
		match, err := regexp.Match(`\([a-z0-9 ]+\)`, []byte(tr))

		if err != nil {
			return err
		}

		if match {

			tr = tr[1 : len(tr)-1]

			splitTerms = strings.Split(tr, " ")

			if len(splitTerms) != 3 {
				return errors.New(
					"invalid transition format for string: " +
						tr +
						", expected exactly 3 terms separated by space")
			}

			newTransition = Transition{splitTerms[0], splitTerms[1], splitTerms[2]}

			fa.Transitions = append(fa.Transitions, newTransition)

		} else {

			return errors.New("invalid transition format for string: " + tr)
		}
	}

	return nil
}

func (fa *FiniteAutomaton) validateTransitions() error {

	for _, tr := range fa.Transitions {

		validTerminal := false
		validCurrentState := false
		validTargetState := false

		for _, term := range fa.Alphabet {

			if tr.NextTerminal == term {
				validTerminal = true
			}
		}

		for _, st := range fa.States {

			if tr.CurrentState == st {
				validCurrentState = true
			}
			if tr.TargetState == st {
				validTargetState = true
			}
		}

		if !validTerminal {

			return errors.New("the 2nd term of the transition must a valid terminal from the alphabet")
		}

		if !validCurrentState || !validTargetState {

			return errors.New("the 1st and 3rd terms of the transition must be valid states")
		}
	}
	return nil
}

func (fa *FiniteAutomaton) validateSequence(sequence string) bool {

	for _, term := range sequence {

		validTerm := false

		for _, alphabetTerm := range fa.Alphabet {

			if string(term) == alphabetTerm {
				validTerm = true
				break
			}
		}

		if !validTerm {
			return false
		}
	}

	return true
}
