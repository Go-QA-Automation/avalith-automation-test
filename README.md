# go-selenium-avalith-web-test


:pushpin: ***Description***

`Simple Navigation testing to Avalith webpage. In this case, we use the next techs:`


:small_blue_diamond: Golang

`An open source programming language that makes it easy to build simple, reliable, and efficient software.`

:small_blue_diamond: Selenium

`Is a free (open-source) automated testing framework used to validate web applications across different browsers and platforms. You can use multiple programming languages like Java, C#, Python etc to create Selenium Test Scripts. Testing done using the Selenium testing tool is usually referred to as Selenium Testing.`

:small_blue_diamond: Cucumber

`Is a testing tool that supports Behavior Driven Development (BDD). It offers a way to write tests that anybody can understand, regardless of their technical knowledge. In BDD, users (business analysts, product owners) first write scenarios or acceptance tests that describe the behavior of the system from the customer's perspective, for review and sign-off by the product owners before developers write their codes.`

:small_blue_diamond: Godog

`Godog is an open source behavior-driven development framework for the Go programming language.`


:small_blue_diamond: Cucumber-html-reporter

`Generate Cucumber HTML reports with pie charts`


:pushpin:  ***How To generate the report?***

```
$godog feature/navigation.feature --format=cucumber > log/report.json
```

:pushpin: ***Installation of 'Cucumber HTML reporter?***

```
> npm install cucumber-html-reporter --save-dev
```

:pushpin: ***Create reporter.js file***

`Create the file and paste the next options. Previously you should modify jsonFile and output to your prefered destination. In this case we have log/report.json and in output log/report.html `

```
var reporter = require('cucumber-html-reporter');

var options = {
        theme: 'bootstrap',
        jsonFile: 'log/report.json',
        output: 'log/report.html', 
        reportSuiteAsScenarios: true,
        scenarioTimestamp: true,
        launchReport: true,
        metadata: {
            "App Version":"0.3.2",
            "Test Environment": "STAGING",
            "Browser": "Chrome  54.0.2840.98",
            "Platform": "Windows 10",
            "Parallel": "Scenarios",
            "Executed": "Remote"
        }
    };

    reporter.generate(options);

```

:pushpin: ***Generate the report and post it on log/report.json***

```> $godog feature/navitagion.feature --format=cucumber > log/report.json ```

:pushpin: ***Open the generated report***

`node reporter.js`

`Now, you should see the generated report` 

![image](https://user-images.githubusercontent.com/32901911/114191517-e2539b80-9922-11eb-8ba2-6ca7bf389cff.png)
