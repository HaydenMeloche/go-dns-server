***Note:*** *This project is for learning how DNS works and how to implement it in GO. It is not intended to be used in production as it always responds with the same IP result.*

Over the holidays I came across this interesting guide on how DNS servers work
https://github.com/EmilHernvall/dnsguide/blob/b52da3b32b27c81e5c6729ac14fe01fef8b1b593/chapter1.md
and decided to implement it in GO to learn more about how DNS works and to play around with GO. 

This project implements most of the DNS specification but hardcodes all answers to any questions to the same IP address.

If you're looking for a DNS server to use in production, this is not it. If you're looking for a DNS server to learn how DNS works, this might be it.

For learning more, I recommend the link above and this RFC:
https://datatracker.ietf.org/doc/html/rfc1035#section-4.1