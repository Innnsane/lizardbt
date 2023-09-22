package conv

func List[In interface{}, Out interface{}](in []*In, convert func(*In) Out) (out []*Out) {
	if len(in) == 0 {
		return nil
	}
	out = make([]*Out, 0, len(in))
	for i := range in {
		h := convert(in[i])
		out = append(out, &h)
	}
	return
}
