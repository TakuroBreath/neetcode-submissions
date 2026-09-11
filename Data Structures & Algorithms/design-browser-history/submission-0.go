type Record struct {
	prev *Record
	site string
	next *Record
}

type BrowserHistory struct {
	rec *Record
}


func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{
		rec: &Record{site: homepage},
	}
}

func (this *BrowserHistory) Visit(url string)  {
	this.rec.next = &Record{site: url, prev: this.rec}
	this.rec = this.rec.next
}


func (this *BrowserHistory) Back(steps int) string {
    for steps > 0 && this.rec.prev != nil {
		this.rec = this.rec.prev
		steps--
	}

	return this.rec.site
}


func (this *BrowserHistory) Forward(steps int) string {
    for steps > 0 && this.rec.next != nil {
		this.rec = this.rec.next
		steps--
	}

	return this.rec.site
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */