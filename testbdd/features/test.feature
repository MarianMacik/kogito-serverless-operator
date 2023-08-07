Feature: Deploy SonataFlow Operator

#  Background:
#    Given Namespace is created
#    And Kogito Operator is deployed
#
#  Scenario Outline: Build <runtime> remote S2I with native <native> using KogitoBuild
#    When Build <runtime> example service "kogito-<runtime>-examples/<example-service>" with configuration:
#      | config | native | <native> |
#
#    Then Build "<example-service>-builder" is complete after <minutes> minutes
#    And Build "<example-service>" is complete after 5 minutes
#    And Kogito Runtime "<example-service>" has 1 pods running within 5 minutes
#    And Service "<example-service>" with process name "orders" is available within 2 minutes
#
#    Examples:
#      |runtime|example-service|native|minutes|
  @wip
  Scenario Outline: Test Scenario Outline
    Given Namespace is created
    When SonataFlow Operator is deployed
    When SonataFlowPlatform is deployed
    When SonataFlow orderprocessing example is deployed
    Then SonataFlow "orderprocessing" has the condition "Running" set to "True" within 1 minutes
#    Then Wait 1 seconds
    Then SonataFlow "orderprocessing" is addressable within 1 minute
    Then HTTP POST request as Cloud Event on SonataFlow "orderprocessing" is successful within 1 minute with path "", headers "content-type= application/json,ce-specversion= 1.0,ce-source= /from/localhost,ce-type= orderEvent,ce-id= f0643c68-609c-48aa-a820-5df423fa4fe0" and body:
    """json
    {"id":"f0643c68-609c-48aa-a820-5df423fa4fe0",
    "country":"Czech Republic",
    "total":10000,
    "description":"iPhone 12"
    }
    """

    Then Deployment "event-listener" pods log contains text 'source: /process/shippinghandling' within 1 minutes
    Then Deployment "event-listener" pods log contains text 'source: /process/fraudhandling' within 1 minutes
    Then Deployment "event-listener" pods log contains text '{"id":"f0643c68-609c-48aa-a820-5df423fa4fe0","country":"Czech Republic","total":10000,"description":"iPhone 12","fraudEvaluation":true,"shipping":"international"}' within 1 minutes
    Then Wait 120 seconds
    #Then SonataFlow Operator has 1 pod running
#    Then Asked to print out the name "<name>"
#    Then Print the name "<name>"

    Examples:
      |name|
    |   Marian    |
    #|             Marek|


  @generation
  Scenario Outline: Test Scenario Outline
#    Then SonataFlow "orderprocessing" has the condition "Running" set to "True" within 1 minutes
#    Then Wait 120 seconds
#    Then SonataFlow "orderprocessing" has the condition "Running" set to "True" within 1 minutes
    Then HTTP POST request as Cloud Event on SonataFlow "orderprocessing" is successful within 1 minute with path "", headers "content-type: application/json,ce-specversion: 1.0,ce-source: /from/localhost,ce-type: orderEvent,ce-id: f0643c68-609c-48aa-a820-5df423fa4fe0" and body:
    """json
    {"id":"f0643c68-609c-48aa-a820-5df423fa4fe0",
    "country":"Brazil",
    "total":10000,
    "description":"iPhone 12"
    }
    """

    Examples:
      |name|
      |   Marian    |
    #|             Marek|
