package main

import (
	"fmt"
	"os"

	aws "github.com/aws/aws-sdk-go/aws"
	session "github.com/aws/aws-sdk-go/aws/session"
	ce "github.com/aws/aws-sdk-go/service/costexplorer"
)

func main() {

	start := os.Args[1] //Must be in YYYY-MM-DD Format
	end := os.Args[2]
	granularity := "MONTHLY"
	metrics := []string{
		"BlendedCost",
		"UsageQuantity",
		// "NormalizedUsageAmount",
		// "AmortizedCost",
		// "NetAmortizedCost",
		// "NetUnblendedCost",
		// "UnblendedCost",
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// Create Cost Explorer Service Client
	svc := ce.New(sess)

	result, err := svc.GetCostAndUsage(&ce.GetCostAndUsageInput{
		TimePeriod: &ce.DateInterval{
			Start: aws.String(start),
			End:   aws.String(end),
		},
		Granularity: aws.String(granularity),
		GroupBy: []*ce.GroupDefinition{
			&ce.GroupDefinition{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("SERVICE"),
			},
		},
		Metrics: aws.StringSlice(metrics),
	})
	if err != nil {
		exitErrorf("Unable to generate usage, %v", err)
	}

	fmt.Println("Cost Report:", result.ResultsByTime)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
