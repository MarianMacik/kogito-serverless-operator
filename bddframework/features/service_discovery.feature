@discovery
Feature: Discovery with onboarding

  Background:
    Given Namespace is created

  Scenario Outline: Deploy onboarding example
    Given Kogito Operator is deployed
    
    When Deploy quarkus example service "onboarding-example/hr" with configuration:
      | config | native                    | <native> |
      | label  | department/first          | process  |
      | label  | id                        | process  |
      | label  | employee-validation/first | process  |

    And Deploy quarkus example service "onboarding-example/payroll" with configuration:
      | config | native         | <native> |
      | label  | taxes/rate     | process  |
      | label  | vacations/days | process  |
      | label  | payments/date  | process  |

    And Deploy quarkus example service "onboarding-example/onboarding" with configuration:
      | config | native     | <native> |
      | label  | onboarding | process  |

    And Kogito application "hr" has 1 pods running within <minutes> minutes
    And Kogito application "payroll" has 1 pods running within <minutes> minutes
    And Kogito application "onboarding" has 1 pods running within <minutes> minutes

    Then HTTP POST request on service "onboarding" is successful within 2 minutes with path "onboarding" and body:
      """json
      { 
        "employee" : {
          "firstName" : "Mark", 
          "lastName" : "Test", 
          "personalId" : "xxx-yy-zzz", 
          "birthDate" : "1995-12-10T14:50:12.123+02:00", 
          "address" : {
            "country" : "US", 
            "city" : "Boston", 
            "street" : "any street 3", 
            "zipCode" : "10001"
          }
        }
      }
      """
    
    Examples: Non Native
      | native   | minutes |
      | disabled | 10      |

    # disabled because of https://issues.redhat.com/browse/KOGITO-1357
    @disabled
    @native
    Examples: Native
      | native  | minutes |
      | enabled | 20      |