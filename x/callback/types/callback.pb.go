... (unchanged content) ...
// QueryCallbacksRequest defines the request structure for the QueryCallbacks RPC method.
message QueryCallbacksRequest {
    int64 block_height = 1; // The block height to query callbacks from
    int32 page = 2;         // The page number for pagination
    int32 limit = 3;        // The number of callbacks to return per page
}
... (unchanged content) ...