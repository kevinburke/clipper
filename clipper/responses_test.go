package clipper

var samplePages = []string{"TRANSACTION HISTORY FOR\nCARD 1202728442\nTRANSACTION TYPE\tLOCATION\tROUTE\tPRODUCT\tDEBIT\tCREDIT\tBALANCE*\n12/12/2017 09:17 AM\tSingle-tag fare payment\tSAM bus\tLOC\tClipper Cash\t2.05\t\t146.85\n12/12/2017 06:21 PM\tSingle-tag fare payment\tSAM bus\tLOC\tClipper Cash\t2.05\t\t144.80\n12/14/2017 09:10 AM\tSingle-tag fare payment\tSAM bus\tLOC\tClipper Cash\t2.05\t\t142.75\n12/14/2017 12:08 PM\tSingle-tag fare payment\tSAM bus\tLOC\tClipper Cash\t2.05\t\t140.70\n12/16/2017 04:59 PM\tDual-tag entry transaction, maximum fare deducted (purse debit)\tBelmont\t\tClipper Cash\t12.20\t\t128.50\n12/16/2017 05:53 PM\tDual-tag exit transaction, fare adjustment (purse rebate)\t4th and King (Caltrain)\t\tClipper Cash\t\t6.75\t135.25\n12/16/2017 11:28 PM\tDual-tag entry transaction, no fare deduction\t16th St Mission\t\tClipper Cash\t\t\t135.25\n12/17/2017 12:23 AM\tDual-tag exit transaction, fare payment\tMillbrae (BART)\t\tClipper Cash\t4.60\t\t130.65\n12/17/2017 12:24 AM\tDual-tag entry transaction, maximum fare deducted (purse debit)\tMillbrae (Caltrain)\t\tClipper Cash\t12.20\t\t118.45\n12/17/2017 12:49 AM\tDual-tag exit transaction, fare adjustment (purse rebate)\tBelmont\t\tClipper Cash\t\t9.00\t127.45\n12/18/2017 08:02 AM\tDual-tag entry transaction, maximum fare deducted (purse debit)\tMillbrae (Caltrain)\t\tClipper Cash\t12.20\t\t115.25\n12/18/2017 08:02 AM\tDual-tag exit transaction, fare adjustment (purse rebate)\tMillbrae (Caltrain)\t\tClipper Cash\t\t12.20\t127.45\n12/18/2017 08:03 AM\tDual-tag entry transaction, no fare deduction\tMillbrae (BART)\t\tClipper Cash\t\t\t127.45\n12/18/2017 08:49 AM\tDual-tag exit transaction, fare payment\tMontgomery (BART)\t\tClipper Cash\t4.65\t\t122.80\n12/18/2017 10:22 AM\tDual-tag entry transaction, no fare deduction\tMontgomery (BART)\t\tClipper Cash\t\t\t122.80\n12/18/2017 11:13 AM\tDual-tag exit transaction, fare payment\tMillbrae (BART)\t\tClipper Cash\t4.65\t\t118.15\n12/18/2017 11:24 AM\tDual-tag entry transaction, maximum fare deducted (purse debit)\tMillbrae (Caltrain)\t\tClipper Cash\t12.20\t\t105.95\n12/18/2017 11:44 AM\tDual-tag exit transaction, fare adjustment (purse rebate)\tBelmont\t\tClipper Cash\t\t9.00\t114.95\n12/20/2017 04:44 PM\tDual-tag entry transaction, maximum fare deducted (purse debit)\tBelmont\t\tClipper Cash\t12.20\t\t102.75\n12/20/2017 05:04 PM\tDual-tag exit transaction, fare adjustment (purse rebate)\tMillbrae (Caltrain)\t\tClipper Cash\t\t9.00\t111.75\n12/20/2017 05:05 PM\tDual-tag entry transaction, no fare deduction\tMillbrae (BART)\t\tClipper Cash\t\t\t111.75\n12/20/2017 05:46 PM\tDual-tag exit transaction, fare payment\tPowell St (BART)\t\tClipper Cash\t4.65\t\t107.10\n12/20/2017 07:05 PM\tSingle-tag fare payment\tPowell (Muni)\tNONE\tClipper Cash\t2.50\t\t104.60\n12/20/2017 11:33 PM\tDual-tag entry transaction, no fare deduction\tPowell St (BART)\t\tClipper Cash\t\t\t104.60\n12/21/2017 12:12 AM\tDual-tag exit transaction, fare payment\tMillbrae (BART)\t\tClipper Cash\t4.65\t\t99.95\n01/05/2018 09:02 AM\tSingle-tag fare payment\tSAM bus\tLOC\tClipper Cash\t2.05\t\t97.90\n01/16/2018 08:36 PM\tSingle-tag fare payment\tSAM bus\tLOC\tClipper Cash\t2.05\t\t95.85\n01/31/2018 08:34 AM\tSingle-tag fare payment\tSAM bus\tLOC\tClipper Cash\t2.05\t\t93.80\n01/31/2018 10:37 AM\tSingle-tag fare payment\tSAM bus\tLOC\tClipper Cash\t2.05\t\t91.75\n02/10/2018\t\t\t\t\t\t\tPage 1 of", "TRANSACTION TYPE\tLOCATION\tROUTE\tPRODUCT\tDEBIT\tCREDIT\tBALANCE*\n02/01/2018 02:08 PM\tDual-tag entry transaction, no fare deduction\tMontgomery (BART)\t\tClipper Cash\t\t\t91.75\n02/01/2018 02:13 PM\tDual-tag exit transaction, fare payment\tCivic Center (BART)\t\tClipper Cash\t2.00\t\t89.75\n* If there is a discrepancy in the listing of the card balance, it may be due to a transaction not reaching the central system. Please contact the Customer Service Center at 877-878-8883 with any questions.\n02/10/2018\t\t\t\t\t\t\tPage 2 of"}