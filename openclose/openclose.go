package stream

type FacebookStream struct {
}

type YoutubeStream struct {
}

type Stream struct {
	StreamType string
}

func (st *Stream) GetSettings() {
	// common code for all streams

	if st.StreamType == "Facebook" {
		// Get settings specific to Facebook
	} else if st.StreamType == "Youtube" {
		// Get settings specific to Youtube
	}

}

//===========================================================================
type Capabilities interface {
	SupportsHD() bool
	SupportsComposition() bool
}

type Stream1 struct {
	capabilities Capabilities
}

func (st *Stream1) GetSettings() {
	// common code for all streams

	if st.capabilities.SupportsHD() {
		// Get settings specific to Facebook
	}

	if st.capabilities.SupportsComposition() {
		// Get settings specific to Youtube
	}

}
