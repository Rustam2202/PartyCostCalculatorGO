<html>

<head>
    <h2>Party Cost Calculator</h2>
    This application distributs debts between the participants of the join event with uneven expenses.
</head>

<body>
    <div>
        <h4>Why it's needed</h4>
        <a>
            If you are celebrating a common event (for example, a corporate party or New Year's Eve) with your friends
            or colleagues, then spending money on food, drinks, etc. is inevitable. Each of the participants may buy
            something specific and spend some amount or may not spend at all. After the event (especially with a
            hangover), it can be quite difficult to calculate who owes whom so that in the end everyone spends the same
            amount. This application should help to cope with this.
        </a>
    </div>
    <div>
        <h4>
            How it's works
        </h4>
        <a>
            As a input values used a JSON-database that contains participant names, spents and count of persons if you 
            spent for youself or some group (like your family member).
            The data is sorted by the amount spent criteria, calculated balance sheet status (overpayment or debt) of
            each participant and using the two iterator method debtors are determined with the indication of the
            recipient and the amount of the debt. Thus, participants with the maximum overpayment and debt are
            calculated first, which allows minimizing the number of payments. You can also specify the threshold for
            rounding decimal places.
        </a>
    </div>
     <div>
        <h4>
            How it's use
        </h4>
        <a>
            Use <code>make</code> commands such as <code>build</code> or <code>run</code> to start application.
            Your POST request must be in JSON format, for example:
            <pre>
{
  "persons": [
    {
      "name": "Participant 1",
      "spent": 800,
      "factor": 2,
    },
    {
      "name": "Participant 2",
      "spent": 600,
    },
    {
     "name": "Participant 3",
    }
  ]
}</pre>
            Enter a <code>name</code> of participant, <code>spent</code> and <code>factor</code> if one participant exist more than one person. By default <code>spent = 0</code> and <code>factor = 1</code>.
            In json response you can take info:
            <pre>
{
  "Id": 0,
  "persons": [
  {
      "Id": 0,
      "name": "Participant 1",
      "spent": 800,
      "factor": 2,
      "Balance": 0,
      "IndeptedTo": null
    },
    {
      "Id": 0,
      "name": "Participant 2",
      "spent": 600,
      "factor": 1,
      "Balance": 0,
      "IndeptedTo": null
    },
    {
      "Id": 0,
      "name": "Participant 3",
      "spent": 0,
      "factor": 1,
      "Balance": 0,
      "IndeptedTo": {
        "Participant 1": 100,
        "Participant 2": 250
      }
    }
  ],
  "AllPersonsCount": 4,
  "AverageAmount": 350,
  "TotalAmount": 1400
}</pre>
        </a>
    </div>
    <div>
        <h4>
            To do
        </h4>
        <a>
            <ul>
                <li></li>
                <li></li>
            </ul>
        </a>
    </div>
</body>

</html>
