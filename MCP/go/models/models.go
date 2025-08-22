package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// PropertiesItem represents the PropertiesItem schema from the OpenAPI specification
type PropertiesItem struct {
	TypeField string `json:"_type"`
	Text string `json:"text,omitempty"` // Text representation of an item.
}

// ImageTagRegion represents the ImageTagRegion schema from the OpenAPI specification
type ImageTagRegion struct {
	Displayrectangle NormalizedQuadrilateral `json:"displayRectangle"` // Defines a region of an image. The region is a convex quadrilateral defined by coordinates of its top left, top right, bottom left, and bottom right points. The coordinates are fractional values of the original image's width and height in the range 0.0 through 1.0.
	Queryrectangle NormalizedQuadrilateral `json:"queryRectangle"` // Defines a region of an image. The region is a convex quadrilateral defined by coordinates of its top left, top right, bottom left, and bottom right points. The coordinates are fractional values of the original image's width and height in the range 0.0 through 1.0.
}

// StructuredValue represents the StructuredValue schema from the OpenAPI specification
type StructuredValue struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
}

// Offer represents the Offer schema from the OpenAPI specification
type Offer struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
}

// Thing represents the Thing schema from the OpenAPI specification
type Thing struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
}

// Query represents the Query schema from the OpenAPI specification
type Query struct {
	Displaytext string `json:"displayText,omitempty"` // The display version of the query term.
	Searchlink string `json:"searchLink,omitempty"` // The URL that you use to get the results of the related search. Before using the URL, you must append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header. Use this URL if you're displaying the results in your own user interface. Otherwise, use the webSearchUrl URL.
	Text string `json:"text"` // The query string. Use this string as the query term in a new search request.
	Thumbnail ImageObject `json:"thumbnail,omitempty"` // Defines an image.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL that takes the user to the Bing search results page for the query.
}

// AggregateOffer represents the AggregateOffer schema from the OpenAPI specification
type AggregateOffer struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Pricecurrency string `json:"priceCurrency,omitempty"` // The monetary currency. For example, USD.
	Seller Organization `json:"seller,omitempty"` // Defines an organization.
	Aggregaterating AggregateRating `json:"aggregateRating,omitempty"` // Defines the metrics that indicate how well an item was rated by others.
	Availability string `json:"availability,omitempty"` // The item's availability. The following are the possible values: Discontinued, InStock, InStoreOnly, LimitedAvailability, OnlineOnly, OutOfStock, PreOrder, SoldOut.
	Lastupdated string `json:"lastUpdated,omitempty"` // The last date that the offer was updated. The date is in the form YYYY-MM-DD.
	Price float32 `json:"price,omitempty"` // The item's price.
}

// Organization represents the Organization schema from the OpenAPI specification
type Organization struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
}

// ImageKnowledge represents the ImageKnowledge schema from the OpenAPI specification
type ImageKnowledge struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
}

// ResponseBase represents the ResponseBase schema from the OpenAPI specification
type ResponseBase struct {
	TypeField string `json:"_type,omitempty"`
}

// Intangible represents the Intangible schema from the OpenAPI specification
type Intangible struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
}

// VisualSearchRequest represents the VisualSearchRequest schema from the OpenAPI specification
type VisualSearchRequest struct {
	Imageinfo ImageInfo `json:"imageInfo,omitempty"` // A JSON object that identities the image to get insights of . It also includes the optional crop area that you use to identify the region of interest in the image.
	Knowledgerequest KnowledgeRequest `json:"knowledgeRequest,omitempty"` // A JSON object containing information about the request, such as filters for the resulting actions.
}

// Filters represents the Filters schema from the OpenAPI specification
type Filters struct {
	Site string `json:"site,omitempty"` // The URL of the site to return similar images and similar products from. (e.g., "www.bing.com", "bing.com").
}

// ImageRecipesAction represents the ImageRecipesAction schema from the OpenAPI specification
type ImageRecipesAction struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Text string `json:"text,omitempty"` // Text content of this creative work.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Displayname string `json:"displayName,omitempty"` // A display name for the action.
	Istopaction bool `json:"isTopAction,omitempty"` // A Boolean representing whether this result is the top action.
	Result []Thing `json:"result,omitempty"` // The result produced in the action.
	Serviceurl string `json:"serviceUrl,omitempty"` // Use this URL to get additional data to determine how to take the appropriate action. For example, the serviceUrl might return JSON along with an image URL.
	TypeField string `json:"_type,omitempty"` // Specifies the sub-class of the action.
	Actiontype string `json:"actionType,omitempty"` // A string representing the type of action.
}

// AggregateRating represents the AggregateRating schema from the OpenAPI specification
type AggregateRating struct {
	TypeField string `json:"_type"`
	Text string `json:"text,omitempty"` // Text representation of an item.
	Ratingvalue float32 `json:"ratingValue"` // The mean (average) rating. The possible values are 1.0 through 5.0.
	Bestrating float32 `json:"bestRating,omitempty"` // The highest rated review. The possible values are 1.0 through 5.0.
}

// Response represents the Response schema from the OpenAPI specification
type Response struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
}

// ImageAction represents the ImageAction schema from the OpenAPI specification
type ImageAction struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"` // Text content of this creative work.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Displayname string `json:"displayName,omitempty"` // A display name for the action.
	Istopaction bool `json:"isTopAction,omitempty"` // A Boolean representing whether this result is the top action.
	Result []Thing `json:"result,omitempty"` // The result produced in the action.
	Serviceurl string `json:"serviceUrl,omitempty"` // Use this URL to get additional data to determine how to take the appropriate action. For example, the serviceUrl might return JSON along with an image URL.
}

// Identifiable represents the Identifiable schema from the OpenAPI specification
type Identifiable struct {
	TypeField string `json:"_type,omitempty"`
}

// Person represents the Person schema from the OpenAPI specification
type Person struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
}

// ImagesModule represents the ImagesModule schema from the OpenAPI specification
type ImagesModule struct {
	Value []ImageObject `json:"value,omitempty"` // A list of images.
}

// ImageObject represents the ImageObject schema from the OpenAPI specification
type ImageObject struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Text string `json:"text,omitempty"` // Text content of this creative work.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Hostpagedisplayurl string `json:"hostPageDisplayUrl,omitempty"` // Display URL of the page that hosts the media object.
	Hostpageurl string `json:"hostPageUrl,omitempty"` // URL of the page that hosts the media object.
	Width int `json:"width,omitempty"` // The width of the media object, in pixels.
	Contentsize string `json:"contentSize,omitempty"` // Size of the media object content. Use format "value unit" (e.g., "1024 B").
	Contenturl string `json:"contentUrl,omitempty"` // Original URL to retrieve the source (file) for the media object (e.g., the source URL for the image).
	Encodingformat string `json:"encodingFormat,omitempty"` // Encoding format (e.g., png, gif, jpeg, etc).
	Height int `json:"height,omitempty"` // The height of the media object, in pixels.
}

// MediaObject represents the MediaObject schema from the OpenAPI specification
type MediaObject struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Text string `json:"text,omitempty"` // Text content of this creative work.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
}

// ImageTag represents the ImageTag schema from the OpenAPI specification
type ImageTag struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Code string `json:"code"` // The error code that identifies the category of error.
	Message string `json:"message"` // A description of the error.
	Moredetails string `json:"moreDetails,omitempty"` // A description that provides additional information about the error.
	Parameter string `json:"parameter,omitempty"` // The parameter in the request that caused the error.
	Subcode string `json:"subCode,omitempty"` // The error code that further helps to identify the error.
	Value string `json:"value,omitempty"` // The parameter's value in the request that was not valid.
}

// ImageRelatedSearchesAction represents the ImageRelatedSearchesAction schema from the OpenAPI specification
type ImageRelatedSearchesAction struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"` // Text content of this creative work.
	Serviceurl string `json:"serviceUrl,omitempty"` // Use this URL to get additional data to determine how to take the appropriate action. For example, the serviceUrl might return JSON along with an image URL.
	Displayname string `json:"displayName,omitempty"` // A display name for the action.
	Istopaction bool `json:"isTopAction,omitempty"` // A Boolean representing whether this result is the top action.
	Result []Thing `json:"result,omitempty"` // The result produced in the action.
	TypeField string `json:"_type,omitempty"` // Specifies the sub-class of the action.
	Actiontype string `json:"actionType,omitempty"` // A string representing the type of action.
}

// ImageShoppingSourcesAction represents the ImageShoppingSourcesAction schema from the OpenAPI specification
type ImageShoppingSourcesAction struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Text string `json:"text,omitempty"` // Text content of this creative work.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Serviceurl string `json:"serviceUrl,omitempty"` // Use this URL to get additional data to determine how to take the appropriate action. For example, the serviceUrl might return JSON along with an image URL.
	Displayname string `json:"displayName,omitempty"` // A display name for the action.
	Istopaction bool `json:"isTopAction,omitempty"` // A Boolean representing whether this result is the top action.
	Result []Thing `json:"result,omitempty"` // The result produced in the action.
	TypeField string `json:"_type,omitempty"` // Specifies the sub-class of the action.
	Actiontype string `json:"actionType,omitempty"` // A string representing the type of action.
}

// RecipesModule represents the RecipesModule schema from the OpenAPI specification
type RecipesModule struct {
	Value []Recipe `json:"value,omitempty"` // A list of recipes.
}

// RelatedSearchesModule represents the RelatedSearchesModule schema from the OpenAPI specification
type RelatedSearchesModule struct {
	Value []Query `json:"value,omitempty"` // A list of related searches.
}

// ImageEntityAction represents the ImageEntityAction schema from the OpenAPI specification
type ImageEntityAction struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Text string `json:"text,omitempty"` // Text content of this creative work.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Displayname string `json:"displayName,omitempty"` // A display name for the action.
	Istopaction bool `json:"isTopAction,omitempty"` // A Boolean representing whether this result is the top action.
	Result []Thing `json:"result,omitempty"` // The result produced in the action.
	Serviceurl string `json:"serviceUrl,omitempty"` // Use this URL to get additional data to determine how to take the appropriate action. For example, the serviceUrl might return JSON along with an image URL.
	TypeField string `json:"_type,omitempty"` // Specifies the sub-class of the action.
	Actiontype string `json:"actionType,omitempty"` // A string representing the type of action.
}

// ImagesImageMetadata represents the ImagesImageMetadata schema from the OpenAPI specification
type ImagesImageMetadata struct {
	Aggregateoffer AggregateOffer `json:"aggregateOffer,omitempty"` // Defines a list of offers from merchants that are related to the image.
	Recipesourcescount int `json:"recipeSourcesCount,omitempty"` // The number of websites that offer recipes of the food seen in the image.
	Shoppingsourcescount int `json:"shoppingSourcesCount,omitempty"` // The number of websites that sell the products seen in the image.
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
}

// CropArea represents the CropArea schema from the OpenAPI specification
type CropArea struct {
	Left float32 `json:"left"` // The left coordinate of the region to be cropped. The coordinate is a fractional value of the original image's width and is measured from the left edge of the image. Specify the coordinate as a value from 0.0 through 1.0.
	Right float32 `json:"right"` // The right coordinate of the region to be cropped. The coordinate is a fractional value of the original image's width and is measured from the left edge of the image. Specify the coordinate as a value from 0.0 through 1.0.
	Top float32 `json:"top"` // The top coordinate of the region to be cropped. The coordinate is a fractional value of the original image's height and is measured from the top edge of the image. Specify the coordinate as a value from 0.0 through 1.0.
	Bottom float32 `json:"bottom"` // The bottom coordinate of the region to be cropped. The coordinate is a fractional value of the original image's height and is measured from the top edge of the image. Specify the coordinate as a value from 0.0 through 1.0.
}

// ImageInfo represents the ImageInfo schema from the OpenAPI specification
type ImageInfo struct {
	Croparea CropArea `json:"cropArea,omitempty"` // A JSON object consisting of coordinates specifying the four corners of a cropped rectangle within the input image.
	Imageinsightstoken string `json:"imageInsightsToken,omitempty"` // An image insights token. To get the insights token, call one of the Image Search APIs (for example, /images/search). In the search results, the [Image](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#image) object's [imageInsightsToken](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#image-imageinsightstoken) field contains the token. The imageInsightsToken and url fields mutually exclusive; do not specify both. Do not specify an insights token if the request includes the image form data.
	Url string `json:"url,omitempty"` // The URL of the input image. The imageInsightsToken and url fields are mutually exclusive; do not specify both. Do not specify the URL if the request includes the image form data.
}

// KnowledgeRequest represents the KnowledgeRequest schema from the OpenAPI specification
type KnowledgeRequest struct {
	Filters Filters `json:"filters,omitempty"` // A key-value object consisting of filters that may be specified to limit the results returned by the API. Current available filters: site.
}

// Rating represents the Rating schema from the OpenAPI specification
type Rating struct {
	TypeField string `json:"_type"`
	Text string `json:"text,omitempty"` // Text representation of an item.
}

// ImageModuleAction represents the ImageModuleAction schema from the OpenAPI specification
type ImageModuleAction struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Text string `json:"text,omitempty"` // Text content of this creative work.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Istopaction bool `json:"isTopAction,omitempty"` // A Boolean representing whether this result is the top action.
	Result []Thing `json:"result,omitempty"` // The result produced in the action.
	Serviceurl string `json:"serviceUrl,omitempty"` // Use this URL to get additional data to determine how to take the appropriate action. For example, the serviceUrl might return JSON along with an image URL.
	Displayname string `json:"displayName,omitempty"` // A display name for the action.
	TypeField string `json:"_type,omitempty"` // Specifies the sub-class of the action.
	Actiontype string `json:"actionType,omitempty"` // A string representing the type of action.
}

// Recipe represents the Recipe schema from the OpenAPI specification
type Recipe struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Text string `json:"text,omitempty"` // Text content of this creative work.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
}

// CreativeWork represents the CreativeWork schema from the OpenAPI specification
type CreativeWork struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
}

// NormalizedQuadrilateral represents the NormalizedQuadrilateral schema from the OpenAPI specification
type NormalizedQuadrilateral struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
}

// Action represents the Action schema from the OpenAPI specification
type Action struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Text string `json:"text,omitempty"` // Text content of this creative work.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
}

// Point2D represents the Point2D schema from the OpenAPI specification
type Point2D struct {
	TypeField string `json:"_type,omitempty"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource. To use the URL, append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image.
}
