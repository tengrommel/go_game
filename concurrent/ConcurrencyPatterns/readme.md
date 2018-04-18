# Error Handling 
> In concurrent programs, error handling can be be difficult to get right.

- Sometimes, we spend so much time thinking about how our various processes will be sharing information and coordinating, 
we forget to consider how they'll gracefully handle errored states.

- When Go eschewed the popular exception model of errors, it made a statement that error handling was important,
and that as we develop our programs, we should give our error paths the same attention we give our algorithms.

- In that spirit, let's take a look at how we do that when working with multiple concurrent processes.

Here we see that the goroutine has been given no choice in the matter.
>It can't simply swallow the error, and so it does the only sensible thing:

 - It can't simply swallow the error, and so it does the only sensible thing:
    > it prints the error and hopes something is paying attention.
 - Don't put your goroutines in this awkward position.
 
 