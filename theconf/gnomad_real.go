func processVariants(mainWG *sync.WaitGroup, reader io.Reader, hashTable *variant.HashTable,
	opt options) {

	defer mainWG.Done()

	variants := make(chan *vcf.Variant)
	invalids := make(chan vcf.InvalidLine)
	keyed := make(chan keyedGnomad)
	accumulator := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(5)

	go parseVCF(&wg, reader, variants, invalids)
	go consumeInvalids(&wg, invalids)
	go parseGnomadFields(&wg, variants, keyed, opt.ParseParallelism)
	go send(&wg, hashTable, opt.ColumnNameModifier, keyed, accumulator, opt.SendParallelism)
	go count(&wg, accumulator, opt.LogBreakpoint)

	wg.Wait()
}
