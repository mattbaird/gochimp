package gochimp

type WorkflowType map[string]interface{}
type TimeToRun map[string]interface{}

// ------------------------------------------------------------------------------------------------
// TimeToRun
// ------------------------------------------------------------------------------------------------

func SendAsap(days []string, sendAsap bool) *TimeToRun {
	return &TimeToRun{
		"days": days,
		"hours": map[string]bool{
			"hours": sendAsap,
		},
	}
}

func SendBetween(days []string, start string, end string) *TimeToRun {
	return &TimeToRun{
		"days": days,
		"hours": map[string]string{
			"start": start,
			"end":   end,
		},
	}
}

func SendAt(days []string, sendAt string) *TimeToRun {
	return &TimeToRun{
		"days": days,
		"hours": map[string]string{
			"send_at": sendAt,
		},
	}
}

// ------------------------------------------------------------------------------------------------
// Workflows
// ------------------------------------------------------------------------------------------------

func WorkflowRecurringEvent(emailCount int, mergeFieldTrigger string, runtime TimeToRun) *WorkflowType {
	return &WorkflowType{
		"workflow_type":         "recurringEvent",
		"runtime":               runtime,
		"workflow_emails_count": emailCount,
		"merge_field_trigger":   mergeFieldTrigger,
	}

}

func WorkflowSpecialEvent(emailCount int, mergeFieldTrigger string, runtime TimeToRun) *WorkflowType {
	return &WorkflowType{
		"workflow_type":         "specialEvent",
		"runtime":               runtime,
		"workflow_emails_count": emailCount,
		"merge_field_trigger":   mergeFieldTrigger,
	}
}

func WorkflowDateAdded(emailCount int, runtime TimeToRun) *WorkflowType {
	return &WorkflowType{
		"workflow_type":         "dateAdded",
		"runtime":               runtime,
		"workflow_emails_count": emailCount,
	}
}

func WorkflowEmailFollowup(emailCount int, triggerOnImport, sendImmediately bool, runtime TimeToRun) *WorkflowType {
	return &WorkflowType{
		"workflow_type":         "emailFollowup",
		"runtime":               runtime,
		"workflow_emails_count": emailCount,
		"send_immediately":      sendImmediately,
		"trigger_on_import":     triggerOnImport,
	}
}

func WorkflowEmailSeries(emailCount int, triggerOnImport, sendImmediately bool, runtime TimeToRun) *WorkflowType {
	return &WorkflowType{
		"workflow_type":         "emailSeries",
		"runtime":               runtime,
		"workflow_emails_count": emailCount,
		"send_immediately":      sendImmediately,
		"trigger_on_import":     triggerOnImport,
	}
}

func WorkflowWelcomeSeries(emailCount int, triggerOnImport, sendImmediately bool, runtime TimeToRun) *WorkflowType {
	return &WorkflowType{
		"workflow_type":         "welcomeSeries",
		"runtime":               runtime,
		"workflow_emails_count": emailCount,
		"send_immediately":      sendImmediately,
		"trigger_on_import":     triggerOnImport,
	}
}

func WorkflowMandrill(emailCount int, sendImmediately bool, mandrillTags []string, runtime TimeToRun) *WorkflowType {
	return &WorkflowType{
		"workflow_type":         "mandrill",
		"runtime":               runtime,
		"workflow_emails_count": emailCount,
		"send_immediately":      sendImmediately,
		"mandrill_tags":         mandrillTags,
	}
}

func WorkflowVisitURL(emailCount int, sendImmediately bool, goalURL string, runtime TimeToRun) *WorkflowType {
	return &WorkflowType{
		"workflow_type":         "visitUrl",
		"runtime":               runtime,
		"workflow_emails_count": emailCount,
		"send_immediately":      sendImmediately,
		"goal_url":              goalURL,
	}
}

func WorkflowBestCustomer(emailCount int, sendImmediately bool, lifetimePurchaseValue float64, purchaseOrders int, runtime TimeToRun) *WorkflowType {
	return &WorkflowType{
		"workflow_type":           "bestCustomer",
		"runtime":                 runtime,
		"workflow_emails_count":   emailCount,
		"send_immediately":        sendImmediately,
		"lifetime_purchase_value": lifetimePurchaseValue,
		"purchase_orders":         purchaseOrders,
	}
}

func WorkflowProductFollowup(emailCount int, sendImmediately bool, productName string, runtime TimeToRun) *WorkflowType {
	return &WorkflowType{
		"workflow_type":         "productFollowup",
		"runtime":               runtime,
		"workflow_emails_count": emailCount,
		"send_immediately":      sendImmediately,
		"product_name":          productName,
	}
}

func WorkflowCategoryFollowup(emailCount int, sendImmediately bool, categoryName string, runtime TimeToRun) *WorkflowType {
	return &WorkflowType{
		"workflow_type":         "categoryFollowup",
		"runtime":               runtime,
		"workflow_emails_count": emailCount,
		"send_immediately":      sendImmediately,
		"category_name":         sendImmediately,
	}
}

func WorkflowPurchaseFollowup(emailCount int, sendImmediately bool, runtime TimeToRun) *WorkflowType {
	return &WorkflowType{
		"workflow_type":         "purchaseFollowup",
		"runtime":               runtime,
		"workflow_emails_count": emailCount,
		"send_immediately":      sendImmediately,
	}
}
func WorkflowAPI(emailCount int, runtime TimeToRun) *WorkflowType {
	return &WorkflowType{
		"workflow_type":         "api",
		"runtime":               runtime,
		"workflow_emails_count": emailCount,
	}
}
func WorkflowGroupAdd(emailCount int, sendImmediately bool, groupID int, runtime TimeToRun) *WorkflowType {
	return &WorkflowType{
		"workflow_type":         "groupAdd",
		"runtime":               runtime,
		"workflow_emails_count": emailCount,
		"send_immediately":      sendImmediately,
		"group_id":              groupID,
	}
}
func WorkflowGroupRemove(emailCount int, sendImmediately bool, groupID int, runtime TimeToRun) *WorkflowType {
	return &WorkflowType{
		"workflow_type":         "groupRemove",
		"runtime":               runtime,
		"workflow_emails_count": emailCount,
		"send_immediately":      sendImmediately,
		"group_id":              groupID,
	}
}
