Feature: Register a shipment
  As a user
  I want to request a new shipment
  So that I can track the shipment status

  Background:
    Given the system is ready to accept shipment registration requests

  Scenario: Successfully register a new shipment
    Given the origin is "New York"
    And the destination is "Los Angeles"
    When a shipment request is submitted
    Then the registration request is successful
    And the shipment state is "created"
