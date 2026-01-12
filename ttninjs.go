package ttninjs

import (
	"fmt"
	"reflect"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// A TT news item as JSON object -- Derived from
// https://www.iptc.org/std/ninjs/ninjs-schema_1.5.json -- (c) Copyright 2023 TT -
// TT Nyhetsbyrån - tt.se - This document is published under the Creative Commons
// Attribution 3.0 license, see  http://creativecommons.org/licenses/by/3.0/.
type Document struct {
	// An object with information about standard, version and schema this instance is
	// valid against. nar:standard, nar:standardversion and xml:schema issue #43.
	// (Added in version 1.3)
	Standard Standard `json:"$standard,omitempty" yaml:"$standard,omitempty" mapstructure:"$standard,omitempty"`

	// Editorial advice to the receiver of the news object. Only in dev so far. Tests
	// with the C-POP project.
	Advice []AdviceElem `json:"advice,omitempty" yaml:"advice,omitempty" mapstructure:"advice,omitempty"`

	// Alternative identifiers of the item. It is up to the individual provider to
	// name and set type on the alternative identifiers they like to use. nar:altId
	// issue #3. (Added in version 1.3)
	Altids *Altids `json:"altids,omitempty" yaml:"altids,omitempty" mapstructure:"altids,omitempty"`

	// $$TT: Individual assignments to produce content connected with one planning
	// item.
	Assignments Assignments `json:"assignments,omitempty" yaml:"assignments,omitempty" mapstructure:"assignments,omitempty"`

	// Content of news objects which are associated with this news object.
	Associations Associations `json:"associations,omitempty" yaml:"associations,omitempty" mapstructure:"associations,omitempty"`

	// $$TT: Object with properties containing data on upcoming events.
	BodyEvent *BodyEvent `json:"body_event,omitempty" yaml:"body_event,omitempty" mapstructure:"body_event,omitempty"`

	// $$TT: The textual content of the news object as HTML5. Only present if type is
	// PUBL or DATA.
	BodyHtml5 string `json:"body_html5,omitempty" yaml:"body_html5,omitempty" mapstructure:"body_html5,omitempty"`

	// $$TT: One or more objects describing the pages in this delivery.
	BodyPages BodyPages `json:"body_pages,omitempty" yaml:"body_pages,omitempty" mapstructure:"body_pages,omitempty"`

	// $$TT: The textual content of the news object as HTML5. Only present if type is
	// PUBL or DATA. See alternative html5 schemas for details. richhtml5 allow more
	// than the older html5 container
	BodyRichhtml5 string `json:"body_richhtml5,omitempty" yaml:"body_richhtml5,omitempty" mapstructure:"body_richhtml5,omitempty"`

	// $$TT: When the news object is some form of sportsresults, table etc the data is
	// delivered as sportsml. Only present if type is PUBL or DATA.
	BodySportsml string `json:"body_sportsml,omitempty" yaml:"body_sportsml,omitempty" mapstructure:"body_sportsml,omitempty"`

	// $$TT: The textual content of the news object as untagged text. Only present if
	// type is PUBL or DATA.
	BodyText string `json:"body_text,omitempty" yaml:"body_text,omitempty" mapstructure:"body_text,omitempty"`

	// The name(s) of the creator(s) of the content
	Byline string `json:"byline,omitempty" yaml:"byline,omitempty" mapstructure:"byline,omitempty"`

	// Holder of one or more byline objects.
	Bylines []BylinesElem `json:"bylines,omitempty" yaml:"bylines,omitempty" mapstructure:"bylines,omitempty"`

	// The total character count in the article excluding figure captions. (Added in
	// version 1.2 according to issue #27.). nar:charcount $$TT: The total character
	// count in the article excluding figure captions.
	Charcount *float64 `json:"charcount,omitempty" yaml:"charcount,omitempty" mapstructure:"charcount,omitempty"`

	// $$TT: String identifier for who receives commission for this object.
	Commissioncode string `json:"commissioncode,omitempty" yaml:"commissioncode,omitempty" mapstructure:"commissioncode,omitempty"`

	// $$TT: When pubstatus is 'commissioned', this field tells who commissioned it.
	Commissionedby []string `json:"commissionedby,omitempty" yaml:"commissionedby,omitempty" mapstructure:"commissionedby,omitempty"`

	// The date and time when the content of this ninjs object was originally created.
	// For example an old photo that is now handled as a ninjs object.
	// nar:contentCreated (Added in 1.4)
	Contentcreated *time.Time `json:"contentcreated,omitempty" yaml:"contentcreated,omitempty" mapstructure:"contentcreated,omitempty"`

	// The person or organisation claiming the intellectual property for the content.
	Copyrightholder string `json:"copyrightholder,omitempty" yaml:"copyrightholder,omitempty" mapstructure:"copyrightholder,omitempty"`

	// Any necessary copyright notice for claiming the intellectual property for the
	// content.
	Copyrightnotice string `json:"copyrightnotice,omitempty" yaml:"copyrightnotice,omitempty" mapstructure:"copyrightnotice,omitempty"`

	// $$TT Used for items that concern a specific date such as events and planning
	// items. Notice that this holds date only, no time. See also datetime.
	Date *SerializableDate `json:"date,omitempty" yaml:"date,omitempty" mapstructure:"date,omitempty"`

	// $$TT For items that concern a specific date and time. See also date.
	Datetime *time.Time `json:"datetime,omitempty" yaml:"datetime,omitempty" mapstructure:"datetime,omitempty"`

	// $$TT: Textual description of the item as text.
	DescriptionText string `json:"description_text,omitempty" yaml:"description_text,omitempty" mapstructure:"description_text,omitempty"`

	// $$TT: TT editorial information. Can be anything from planned re-relases of
	// object to restrictions. (DEPRECATED, use ednote instead!)
	DescriptionUsage string `json:"description_usage,omitempty" yaml:"description_usage,omitempty" mapstructure:"description_usage,omitempty"`

	// A note that is intended to be read by internal staff at the receiving
	// organisation, but not published to the end-user. (Added in version 1.2 from
	// issue #6.) . ednote: nar:edNote  $$TT: TT will start using ednote and deprecate
	// description_usage
	Ednote string `json:"ednote,omitempty" yaml:"ednote,omitempty" mapstructure:"ednote,omitempty"`

	// The date and time before which all versions of the object are embargoed. If
	// absent, this object is not embargoed.
	Embargoed *time.Time `json:"embargoed,omitempty" yaml:"embargoed,omitempty" mapstructure:"embargoed,omitempty"`

	// $$TT: Textual description of why article is embargoed.
	Embargoedreason string `json:"embargoedreason,omitempty" yaml:"embargoedreason,omitempty" mapstructure:"embargoedreason,omitempty"`

	// $$TT Used for items that concern a specific date such as events and planning
	// items and has a specific enddate. Notice that this holds date only, no time.
	// See also enddatetime.
	Enddate *SerializableDate `json:"enddate,omitempty" yaml:"enddate,omitempty" mapstructure:"enddate,omitempty"`

	// $$TT For items that concern a specific enddate and time. See also enddate.
	Enddatetime *time.Time `json:"enddatetime,omitempty" yaml:"enddatetime,omitempty" mapstructure:"enddatetime,omitempty"`

	// Something which happens in a planned or unplanned manner. nar:?
	Event []EventElem `json:"event,omitempty" yaml:"event,omitempty" mapstructure:"event,omitempty"`

	// The date and time after which the Item is no longer considered editorially
	// relevant by its provider. nar:expires (Added in 1.4)
	Expires *time.Time `json:"expires,omitempty" yaml:"expires,omitempty" mapstructure:"expires,omitempty"`

	// Indicates when the first version of the item was created. (Added in version 1.2
	// from issue #5). nar:firstCreated
	Firstcreated *time.Time `json:"firstcreated,omitempty" yaml:"firstcreated,omitempty" mapstructure:"firstcreated,omitempty"`

	// $$TT: A storytag that this item belong to. A sort of grouping name that can be
	// used over time for a running story. Broader than a slugline, narrower than a
	// media topic.
	Fixture []FixtureElem `json:"fixture,omitempty" yaml:"fixture,omitempty" mapstructure:"fixture,omitempty"`

	// A nature, intellectual or journalistic form of the content. nar:genre. (Added
	// in version 1.3)  $$TT: TT will move sector to genre and deprecate sector.
	Genre []GenreElem `json:"genre,omitempty" yaml:"genre,omitempty" mapstructure:"genre,omitempty"`

	// A brief and snappy introduction to the content, designed to catch the reader's
	// attention
	Headline string `json:"headline,omitempty" yaml:"headline,omitempty" mapstructure:"headline,omitempty"`

	// A party (person or organisation) which originated, modified, enhanced,
	// distributed, aggregated or supplied the content or provided some information
	// used to create or enhance the content. (Added in version 1.2 according to issue
	// #15.) .    infosource:  nar:infoSource
	Infosource []InfosourceElem `json:"infosource,omitempty" yaml:"infosource,omitempty" mapstructure:"infosource,omitempty"`

	// $$TT: Identifier of a grouping job this item belongs to. Typically the id of
	// the job the article belong to, normally something like 327890.
	Job string `json:"job,omitempty" yaml:"job,omitempty" mapstructure:"job,omitempty"`

	// The human language used by the content. The value should follow IETF BCP47
	Language string `json:"language,omitempty" yaml:"language,omitempty" mapstructure:"language,omitempty"`

	// The name of the location from which the content originates.
	Located string `json:"located,omitempty" yaml:"located,omitempty" mapstructure:"located,omitempty"`

	// A MIME type which applies to this object
	Mimetype string `json:"mimetype,omitempty" yaml:"mimetype,omitempty" mapstructure:"mimetype,omitempty"`

	// $TT: TT managed editorial sort order. Priority numbers range from 6 (most
	// important) to 1 (least).
	Newsvalue *int `json:"newsvalue,omitempty" yaml:"newsvalue,omitempty" mapstructure:"newsvalue,omitempty"`

	// Something material, excluding persons. nar:subject
	Object []ObjectElem `json:"object,omitempty" yaml:"object,omitempty" mapstructure:"object,omitempty"`

	// An administrative and functional structure which may act as as a business, as a
	// political party or not-for-profit party. nar:subject
	Organisation []OrganisationElem `json:"organisation,omitempty" yaml:"organisation,omitempty" mapstructure:"organisation,omitempty"`

	// $$TT: Identifier in the originating system/source. DEPRECATED: Will be handled
	// as an altid
	Originaltransmissionreference string `json:"originaltransmissionreference,omitempty" yaml:"originaltransmissionreference,omitempty" mapstructure:"originaltransmissionreference,omitempty"`

	// An individual human being
	Person []PersonElem `json:"person,omitempty" yaml:"person,omitempty" mapstructure:"person,omitempty"`

	// A named location
	Place []PlaceElem `json:"place,omitempty" yaml:"place,omitempty" mapstructure:"place,omitempty"`

	// $$TT: TT Product classification codes. See http://tt.se/spec/product/1.0/
	Product []ProductElem `json:"product,omitempty" yaml:"product,omitempty" mapstructure:"product,omitempty"`

	// An identifier for the structure of the news object. This can be any string but
	// we suggest something identifying the structure of the content such as
	// 'text-only' or 'text-photo'. Profiles are typically provider-specific.
	// nar:profile $$TT: Possible values are PUBL, DATA, INFO or RAW. PUBL is a news
	// item that can be published. DATA is data such as tables and figures (that are
	// not meant to be edited). INFO is for information purposes only (not to be
	// published). RAW is raw data, such as unedited videos, that is meant to be
	// further edited before publishing.
	Profile *Profile `json:"profile,omitempty" yaml:"profile,omitempty" mapstructure:"profile,omitempty"`

	// The publishing status of the news object, its value is *usable* by default.
	// Please note that for information about events that have been canceled the
	// pubstatus of the ttninjs object will still be usable. The cancel information
	// can be found in body_event. $$TT: replaced and comissioned added by TT.
	Pubstatus Pubstatus `json:"pubstatus,omitempty" yaml:"pubstatus,omitempty" mapstructure:"pubstatus,omitempty"`

	// Wrapper for different renditions of non-textual content of the news object
	Renditions Renditions `json:"renditions,omitempty" yaml:"renditions,omitempty" mapstructure:"renditions,omitempty"`

	// $$TT: The identifier of the news object this one is replaced by.
	Replacedby string `json:"replacedby,omitempty" yaml:"replacedby,omitempty" mapstructure:"replacedby,omitempty"`

	// $$TT: Array of identifiers of news objects this object is replacing.
	Replacing []string `json:"replacing,omitempty" yaml:"replacing,omitempty" mapstructure:"replacing,omitempty"`

	// Indicates how complete this representation of a news item is. $$TT: associated
	// is a TT-extension used when the news item appears as an association considered
	// as a link without renditions.
	Representationtype *Representationtype `json:"representationtype,omitempty" yaml:"representationtype,omitempty" mapstructure:"representationtype,omitempty"`

	// $$TT: Array of previous versions of this news object. See
	// http://spec.tt.se/revisions.html
	Revisions []RevisionsElem `json:"revisions,omitempty" yaml:"revisions,omitempty" mapstructure:"revisions,omitempty"`

	// Expression of rights to be applied to content. nar:rightsInfo (Added in 1.4)
	Rightsinfo *Rightsinfo `json:"rightsinfo,omitempty" yaml:"rightsinfo,omitempty" mapstructure:"rightsinfo,omitempty"`

	// $$TT: Designator for the major ways of grouping content (inrikes, utrikes, etc)
	// and PRM for press releases. Not mandatory, often omitted. DEPRECATED and moved
	// to genre.
	Sector *Sector `json:"sector,omitempty" yaml:"sector,omitempty" mapstructure:"sector,omitempty"`

	// $$TT: signals is suggested by AP but not yet included in ninjs. When included
	// it will probably hold a large number of properties.
	Signals Signals `json:"signals,omitempty" yaml:"signals,omitempty" mapstructure:"signals,omitempty"`

	// $$TT: Short name given to article while in production. (DEPRECTED, use slugline
	// instead.)
	Slug string `json:"slug,omitempty" yaml:"slug,omitempty" mapstructure:"slug,omitempty"`

	// A human-readable identifier for the item. (Added in version 1.2 from issue
	// #4.). nar:slugline  $$TT: TT will use slugline and deprecate slug.
	Slugline string `json:"slugline,omitempty" yaml:"slugline,omitempty" mapstructure:"slugline,omitempty"`

	// $$TT: String identifier for originating source of content.
	Source string `json:"source,omitempty" yaml:"source,omitempty" mapstructure:"source,omitempty"`

	// A concept with a relationship to the content. $$TT: Used for content
	// classification in swedish equivalent of IPTC Subject Reference see
	// http://tt.se/spec/subref/1.0/ etc
	Subject []SubjectElem `json:"subject,omitempty" yaml:"subject,omitempty" mapstructure:"subject,omitempty"`

	// A short natural-language name for the item. (Added in version 1.2 according to
	// issue #9). nar:itemMeta/title
	Title string `json:"title,omitempty" yaml:"title,omitempty" mapstructure:"title,omitempty"`

	// An array of objects to allow links to documents about trust indicators.
	// (nar:link) issue #44. (Added in version 1.3)
	Trustindicator []TrustindicatorElem `json:"trustindicator,omitempty" yaml:"trustindicator,omitempty" mapstructure:"trustindicator,omitempty"`

	// The generic news type of this news object. $$TT: TT  added event for items with
	// data describing a coming event.
	Type Type `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`

	// The editorial urgency of the content from 1 to 9. 1 represents the highest
	// urgency, 9 the lowest. $$TT: 1 is most urgent. 4 is normal. Definition here
	// http://tt.se/spec/prio/1.0
	Urgency int `json:"urgency,omitempty" yaml:"urgency,omitempty" mapstructure:"urgency,omitempty"`

	// The identifier for this object
	Uri string `json:"uri" yaml:"uri" mapstructure:"uri"`

	// A natural-language statement about the usage terms pertaining to the content.
	// $$TT: Specifically contains image usage restrictions from TT's suppliers.
	Usageterms string `json:"usageterms,omitempty" yaml:"usageterms,omitempty" mapstructure:"usageterms,omitempty"`

	// The version of the object which is identified by the uri property
	Version string `json:"version,omitempty" yaml:"version,omitempty" mapstructure:"version,omitempty"`

	// The date and time when this version of the object was created
	Versioncreated time.Time `json:"versioncreated,omitempty" yaml:"versioncreated,omitempty" mapstructure:"versioncreated,omitempty"`

	// $$TT: The date and time when this version of the object was persisted. For a
	// photo, versioncreated is when photo was taken, versionstored is when we indexed
	// it to the database.
	Versionstored *time.Time `json:"versionstored,omitempty" yaml:"versionstored,omitempty" mapstructure:"versionstored,omitempty"`

	// $TT: TT managed editorial sort order. Priority numbers range from 1 (most
	// important) to 3 (least). A 0 indicates that the item needs manual attention
	// before publishning. Definitions and sort logic are defined here
	// http://tt.se/spec/webprio/1.0
	Webprio int `json:"webprio,omitempty" yaml:"webprio,omitempty" mapstructure:"webprio,omitempty"`

	// $$TT: The number of the week the item is planned to be published. Mainly used
	// for feature-articles and ready pages. Also showing the week-number of planning
	// and events.
	Week int `json:"week,omitempty" yaml:"week,omitempty" mapstructure:"week,omitempty"`

	// The total number of words in the article excluding figure captions. (Added in
	// version 1.2 according to issue #27.). nar:wordcount
	Wordcount int `json:"wordcount,omitempty" yaml:"wordcount,omitempty" mapstructure:"wordcount,omitempty"`
}

// One advice item
type AdviceElem struct {
	// Environment corresponds to the JSON schema field "environment".
	Environment []AdviceElemEnvironmentElem `json:"environment,omitempty" yaml:"environment,omitempty" mapstructure:"environment,omitempty"`

	// Advice regarding the importance of the content from an emotional perspective.
	// Experimental vocabulary, part of the C-POP project.
	Importance *AdviceElemImportance `json:"importance,omitempty" yaml:"importance,omitempty" mapstructure:"importance,omitempty"`

	// Advice regarding the length of time that the content is considered to be
	// relevant. Experimental vocabulary, part of the C-POP project.
	Lifetime *AdviceElemLifetime `json:"lifetime,omitempty" yaml:"lifetime,omitempty" mapstructure:"lifetime,omitempty"`

	// Role of this advice.
	Role AdviceElemRole `json:"role,omitempty" yaml:"role,omitempty" mapstructure:"role,omitempty"`
}

type AdviceElemEnvironmentElem struct {
	// Code corresponds to the JSON schema field "code".
	Code string `json:"code,omitempty" yaml:"code,omitempty" mapstructure:"code,omitempty"`

	// Scheme corresponds to the JSON schema field "scheme".
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty" mapstructure:"scheme,omitempty"`
}

// Advice regarding the importance of the content from an emotional perspective.
// Experimental vocabulary, part of the C-POP project.
type AdviceElemImportance struct {
	// Present values are: essential, useful and entertaining.
	Code string `json:"code,omitempty" yaml:"code,omitempty" mapstructure:"code,omitempty"`

	// http://cv.iptc.org/newscodes/advice-importance
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty" mapstructure:"scheme,omitempty"`
}

// Advice regarding the length of time that the content is considered to be
// relevant. Experimental vocabulary, part of the C-POP project.
type AdviceElemLifetime struct {
	// Present values are: short, medium, long and evergreen.
	Code string `json:"code,omitempty" yaml:"code,omitempty" mapstructure:"code,omitempty"`

	// http://cv.iptc.org/newscodes/advice-lifetime
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty" mapstructure:"scheme,omitempty"`
}

type AdviceElemRole string

const AdviceElemRolePublish AdviceElemRole = "publish"

// Alternative identifiers of the item. It is up to the individual provider to name
// and set type on the alternative identifiers they like to use. nar:altId issue
// #3. (Added in version 1.3)
type Altids struct {
	// $$TT: Identifier in the originating system/source. TT will move
	// originaltransmissionreference here.
	Originaltransmissionreference string `json:"originaltransmissionreference,omitempty" yaml:"originaltransmissionreference,omitempty" mapstructure:"originaltransmissionreference,omitempty"`
}

// $$TT: Individual assignments to produce content connected with one planning
// item.
type Assignments map[string]Document

// Content of news objects which are associated with this news object.
type Associations map[string]Document

// $$TT: Object with properties containing data on upcoming events.
type BodyEvent struct {
	// $$TT: Information about how to get accreditation to the event.
	Accreditation string `json:"accreditation,omitempty" yaml:"accreditation,omitempty" mapstructure:"accreditation,omitempty"`

	// $$TT: Address to the place where the event will take place.
	Address string `json:"address,omitempty" yaml:"address,omitempty" mapstructure:"address,omitempty"`

	// $$TT: Name of the arena where the event will take place.
	Arena string `json:"arena,omitempty" yaml:"arena,omitempty" mapstructure:"arena,omitempty"`

	// $$TT: Initials of the person doing the last update to the item.
	Changedby string `json:"changedby,omitempty" yaml:"changedby,omitempty" mapstructure:"changedby,omitempty"`

	// $$TT: When the item was last updated in the TT event database.
	Changeddate *time.Time `json:"changeddate,omitempty" yaml:"changeddate,omitempty" mapstructure:"changeddate,omitempty"`

	// $$TT: Name of the city where the event will take place.
	City string `json:"city,omitempty" yaml:"city,omitempty" mapstructure:"city,omitempty"`

	// $$TT: Three letter code for the country where the event will take place.
	Country string `json:"country,omitempty" yaml:"country,omitempty" mapstructure:"country,omitempty"`

	// $$TT: If the event is a trial this property hold the casenumber.
	Courtcasenumber string `json:"courtcasenumber,omitempty" yaml:"courtcasenumber,omitempty" mapstructure:"courtcasenumber,omitempty"`

	// $$TT: Initials of the person creating the item in the TT event database.
	Createdby string `json:"createdby,omitempty" yaml:"createdby,omitempty" mapstructure:"createdby,omitempty"`

	// $$TT: When the item was created in the TT event database.
	Createddate *time.Time `json:"createddate,omitempty" yaml:"createddate,omitempty" mapstructure:"createddate,omitempty"`

	// $$TT: Phone number to call for more information about the event.
	Eventphone string `json:"eventphone,omitempty" yaml:"eventphone,omitempty" mapstructure:"eventphone,omitempty"`

	// $$TT: Status code for the event. Value is normally 1. Canceled events will have
	// 4.
	Eventstatus string `json:"eventstatus,omitempty" yaml:"eventstatus,omitempty" mapstructure:"eventstatus,omitempty"`

	// $$TT: Status for the event as a phrase. Normally 'Planerat'. Canceled events
	// will have 'Inställt'.
	EventstatusText string `json:"eventstatus_text,omitempty" yaml:"eventstatus_text,omitempty" mapstructure:"eventstatus_text,omitempty"`

	// $$TT: Tags of the event.
	Eventtags string `json:"eventtags,omitempty" yaml:"eventtags,omitempty" mapstructure:"eventtags,omitempty"`

	// $$TT: Code for type of event.
	Eventtype string `json:"eventtype,omitempty" yaml:"eventtype,omitempty" mapstructure:"eventtype,omitempty"`

	// $$TT: Type of event as text.
	EventtypeText string `json:"eventtype_text,omitempty" yaml:"eventtype_text,omitempty" mapstructure:"eventtype_text,omitempty"`

	// $$TT: URL to a web site with information about the event.
	Eventurl string `json:"eventurl,omitempty" yaml:"eventurl,omitempty" mapstructure:"eventurl,omitempty"`

	// $$TT: Details on following the event online
	Eventweb string `json:"eventweb,omitempty" yaml:"eventweb,omitempty" mapstructure:"eventweb,omitempty"`

	// $$TT: If there are more information concerning the event.
	Extraurl string `json:"extraurl,omitempty" yaml:"extraurl,omitempty" mapstructure:"extraurl,omitempty"`

	// $$TT: For events in Sweden, the code of the municipality.
	Municipality string `json:"municipality,omitempty" yaml:"municipality,omitempty" mapstructure:"municipality,omitempty"`

	// $$TT: For events in Sweden the name of the municipality.
	MunicipalityText string `json:"municipality_text,omitempty" yaml:"municipality_text,omitempty" mapstructure:"municipality_text,omitempty"`

	// $$TT: Extra information about the event.
	NoteExtra string `json:"note_extra,omitempty" yaml:"note_extra,omitempty" mapstructure:"note_extra,omitempty"`

	// $$TT: Text intended to be used by TT on planning lists of upcoming events.
	NotePm string `json:"note_pm,omitempty" yaml:"note_pm,omitempty" mapstructure:"note_pm,omitempty"`

	// $$TT: Name of the organizer of the event
	Organizer string `json:"organizer,omitempty" yaml:"organizer,omitempty" mapstructure:"organizer,omitempty"`

	// $$TT: Adress of the organizer of the event
	Organizeraddress string `json:"organizeraddress,omitempty" yaml:"organizeraddress,omitempty" mapstructure:"organizeraddress,omitempty"`

	// $$TT: City name of the organizer of the event
	Organizercity string `json:"organizercity,omitempty" yaml:"organizercity,omitempty" mapstructure:"organizercity,omitempty"`

	// $$TT: Country of the organizer of the event
	Organizercountry string `json:"organizercountry,omitempty" yaml:"organizercountry,omitempty" mapstructure:"organizercountry,omitempty"`

	// $$TT: Mail address to the organizer of the event.
	Organizermail string `json:"organizermail,omitempty" yaml:"organizermail,omitempty" mapstructure:"organizermail,omitempty"`

	// $$TT: Phone number to the organizer of the event.
	Organizerphone string `json:"organizerphone,omitempty" yaml:"organizerphone,omitempty" mapstructure:"organizerphone,omitempty"`

	// $$TT: URL to a web page for the organizer
	Organizerurl string `json:"organizerurl,omitempty" yaml:"organizerurl,omitempty" mapstructure:"organizerurl,omitempty"`

	// $$TT: For events in Sweden, the code of the region.
	Region string `json:"region,omitempty" yaml:"region,omitempty" mapstructure:"region,omitempty"`

	// $$TT: For events in Sweden, the name of the region.
	RegionText string `json:"region_text,omitempty" yaml:"region_text,omitempty" mapstructure:"region_text,omitempty"`
}

// $$TT: One or more objects describing the pages in this delivery.
type BodyPages map[string]interface{}

type BylinesElem struct {
	// The affiliation of the person. Example: SvD/TT
	Affiliation string `json:"affiliation,omitempty" yaml:"affiliation,omitempty" mapstructure:"affiliation,omitempty"`

	// When the complete byline is sent as one string. Same as byline on root level.
	// Example: Albert Jonsson/SvD/TT
	Byline string `json:"byline,omitempty" yaml:"byline,omitempty" mapstructure:"byline,omitempty"`

	// Email address of the person in this byline. albert.jonsson@acme.com
	Email string `json:"email,omitempty" yaml:"email,omitempty" mapstructure:"email,omitempty"`

	// When byline is divided, holds the first name of the person. Example: Albert
	Firstname string `json:"firstname,omitempty" yaml:"firstname,omitempty" mapstructure:"firstname,omitempty"`

	// Initials of byline. Mainly used for records marked as internal. Example: mag
	Initials string `json:"initials,omitempty" yaml:"initials,omitempty" mapstructure:"initials,omitempty"`

	// Whether byline is for internal purposes. Example: true. If not present it means
	// false.
	Internal string `json:"internal,omitempty" yaml:"internal,omitempty" mapstructure:"internal,omitempty"`

	// Jobtitle can differ from role and is normally more connected to the person and
	// not to the combination person-newsItem. Example: Editor in Chief
	Jobtitle string `json:"jobtitle,omitempty" yaml:"jobtitle,omitempty" mapstructure:"jobtitle,omitempty"`

	// When byline is divided, holds the last name of the person. Example: Jonsson
	Lastname string `json:"lastname,omitempty" yaml:"lastname,omitempty" mapstructure:"lastname,omitempty"`

	// Phone number of the person in this byline. Example: +46555123456
	Phone string `json:"phone,omitempty" yaml:"phone,omitempty" mapstructure:"phone,omitempty"`

	// Role of the person in the byline in relation to this ttninjs item, as string.
	// Example: Photographer
	Role string `json:"role,omitempty" yaml:"role,omitempty" mapstructure:"role,omitempty"`
}

type ContactinfoType struct {
	// Type would be method of communication like phone, mobile, address
	// etc.
	Type string `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	// Role refers to type and could be private, office et.c.
	Role string `json:"role,omitempty" yaml:"role,omitempty" mapstructure:"role,omitempty"`
	// If this contactinfo object need to be qualified with what language it
	// is in. The value should follow IETF BCP47.
	Lang string `json:"lang,omitempty" yaml:"lang,omitempty" mapstructure:"lang,omitempty"`
	// Human readable name of the contact method, like name for a web page,
	// name of persons twitter account et.c.
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`
	// Actual phone number, email address, web url etc.
	Value   string  `json:"value,omitempty" yaml:"value,omitempty" mapstructure:"value,omitempty"`
	Address Address `json:"address,omitempty" yaml:"address,omitempty" mapstructure:"address,omitempty"`
}

type Address struct {
	// An array of lines to construct an address. The order is important to
	// construct a correct address.
	Lines      []string `json:"lines,omitempty" yaml:"lines,omitempty" mapstructure:"lines,omitempty"`
	Locality   string   `json:"locality,omitempty" yaml:"locality,omitempty" mapstructure:"locality,omitempty"`
	Area       string   `json:"area,omitempty" yaml:"area,omitempty" mapstructure:"area,omitempty"`
	Postalcode string   `json:"postalcode,omitempty" yaml:"postalcode,omitempty" mapstructure:"postalcode,omitempty"`
	Country    string   `json:"country,omitempty" yaml:"country,omitempty" mapstructure:"country,omitempty"`
}

type EventElem struct {
	// The code for the event in a scheme (= controlled vocabulary) which is
	// identified by the scheme property
	Code string `json:"code,omitempty" yaml:"code,omitempty" mapstructure:"code,omitempty"`

	// The name of the event
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// The relationship of the content of the news object to the event
	Rel string `json:"rel,omitempty" yaml:"rel,omitempty" mapstructure:"rel,omitempty"`

	// The identifier of a scheme (= controlled vocabulary) which includes a code for
	// the event
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty" mapstructure:"scheme,omitempty"`
}

type FixtureElem struct {
	// The code for the story in a scheme (= controlled vocabulary) which is
	// identified by the scheme property.
	Code string `json:"code,omitempty" yaml:"code,omitempty" mapstructure:"code,omitempty"`

	// The name of the storytag
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// The relationship of the content to the fixture
	Rel string `json:"rel,omitempty" yaml:"rel,omitempty" mapstructure:"rel,omitempty"`

	// The identifier of a scheme (= controlled vocabulary) which includes a code for
	// the subject. $$TT: http://tt.se/spec/story/1.0/
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty" mapstructure:"scheme,omitempty"`
}

type GenreElem struct {
	// The code for the genre in a scheme (= controlled vocabulary) which is
	// identified by the scheme property
	Code string `json:"code,omitempty" yaml:"code,omitempty" mapstructure:"code,omitempty"`

	// The name of the genre
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// The identifier of a scheme (= controlled vocabulary) which includes a code for
	// the genre. Normally  http://cv.iptc.org/newscodes/genre/
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty" mapstructure:"scheme,omitempty"`
}

type InfosourceElem struct {
	// The code for the infosource in a scheme (= controlled vocabulary) which is
	// identified by the scheme property
	Code string `json:"code,omitempty" yaml:"code,omitempty" mapstructure:"code,omitempty"`

	// Contactinfo corresponds to the JSON schema field "contactinfo".
	Contactinfo []ContactinfoType `json:"contactinfo,omitempty" yaml:"contactinfo,omitempty" mapstructure:"contactinfo,omitempty"`

	// The name of the infosource
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// The relationship of the content of the news object to the infosource
	Rel string `json:"rel,omitempty" yaml:"rel,omitempty" mapstructure:"rel,omitempty"`

	// The identifier of a scheme (= controlled vocabulary) which includes a code for
	// the infosource
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty" mapstructure:"scheme,omitempty"`
}

type ObjectElem struct {
	// The code for the object in a scheme (= controlled vocabulary) which is
	// identified by the scheme property
	Code string `json:"code,omitempty" yaml:"code,omitempty" mapstructure:"code,omitempty"`

	// The name of the object
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// The relationship of the content of the news object to the object
	Rel string `json:"rel,omitempty" yaml:"rel,omitempty" mapstructure:"rel,omitempty"`

	// The identifier of a scheme (= controlled vocabulary) which includes a code for
	// the object
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty" mapstructure:"scheme,omitempty"`
}

type OrganisationElem struct {
	// The code for the organisation in a scheme (= controlled vocabulary) which is
	// identified by the scheme property
	Code string `json:"code,omitempty" yaml:"code,omitempty" mapstructure:"code,omitempty"`

	// Contactinfo corresponds to the JSON schema field "contactinfo".
	Contactinfo []ContactinfoType `json:"contactinfo,omitempty" yaml:"contactinfo,omitempty" mapstructure:"contactinfo,omitempty"`

	// The name of the organisation
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// The relationship of the content of the news object to the organisation
	Rel string `json:"rel,omitempty" yaml:"rel,omitempty" mapstructure:"rel,omitempty"`

	// The identifier of a scheme (= controlled vocabulary) which includes a code for
	// the organisation
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty" mapstructure:"scheme,omitempty"`

	// Symbols used for a financial instrument linked to the organisation at a
	// specific market place
	Symbols []OrganisationElemSymbolsElem `json:"symbols,omitempty" yaml:"symbols,omitempty" mapstructure:"symbols,omitempty"`
}

type OrganisationElemSymbolsElem struct {
	// Identifier for the marketplace which uses the ticker symbols of the ticker
	// property
	Exchange string `json:"exchange,omitempty" yaml:"exchange,omitempty" mapstructure:"exchange,omitempty"`

	// Compare with hasInstrument in NewsML-G2. Same as symbol in G2.
	Symbol string `json:"symbol,omitempty" yaml:"symbol,omitempty" mapstructure:"symbol,omitempty"`

	// https://cv.iptc.org/newscodes/financialinstrumentsymboltype. Same as type in
	// G2.
	Symboltype string `json:"symboltype,omitempty" yaml:"symboltype,omitempty" mapstructure:"symboltype,omitempty"`

	// Ticker symbol used for the financial instrument
	Ticker string `json:"ticker,omitempty" yaml:"ticker,omitempty" mapstructure:"ticker,omitempty"`
}

type PersonElem struct {
	// The code for the person in a scheme (= controlled vocabulary) which is
	// identified by the scheme property. $$TT: http://tt.se/spec/person/1.0/
	Code string `json:"code,omitempty" yaml:"code,omitempty" mapstructure:"code,omitempty"`

	// Contactinfo corresponds to the JSON schema field "contactinfo".
	Contactinfo []ContactinfoType `json:"contactinfo,omitempty" yaml:"contactinfo,omitempty" mapstructure:"contactinfo,omitempty"`

	// The name of a person
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// The relationship of the content of the news object to the person
	Rel string `json:"rel,omitempty" yaml:"rel,omitempty" mapstructure:"rel,omitempty"`

	// The identifier of a scheme (= controlled vocabulary) which includes a code for
	// the person
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty" mapstructure:"scheme,omitempty"`
}

type PlaceElem struct {
	// The code for the place in a scheme (= controlled vocabulary) which is
	// identified by the scheme property
	Code string `json:"code,omitempty" yaml:"code,omitempty" mapstructure:"code,omitempty"`

	// Contactinfo corresponds to the JSON schema field "contactinfo".
	Contactinfo []ContactinfoType `json:"contactinfo,omitempty" yaml:"contactinfo,omitempty" mapstructure:"contactinfo,omitempty"`

	// $$TT: An optional GeoJSON description of the place.
	GeometryGeojson *PlaceElemGeometryGeojson `json:"geometry_geojson,omitempty" yaml:"geometry_geojson,omitempty" mapstructure:"geometry_geojson,omitempty"`

	// The name of the place
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// The relationship of the content of the news object to the place. $$TT: We use
	// the values land, län, landskap, kommun, ort, delstat, capital and city to
	// indicate the type of area pointed to by the coordinates. Other types can be
	// added.
	Rel string `json:"rel,omitempty" yaml:"rel,omitempty" mapstructure:"rel,omitempty"`

	// The identifier of a scheme (= controlled vocabulary) which includes a code for
	// the place. $$TT: http://tt.se/spec/place/1.0/
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty" mapstructure:"scheme,omitempty"`
}

// $$TT: An optional GeoJSON description of the place.
type PlaceElemGeometryGeojson struct {
	// Array of coordinate pairs, but in our case on pair.
	Coordinates []float64 `json:"coordinates,omitempty" yaml:"coordinates,omitempty" mapstructure:"coordinates,omitempty"`

	// What type of coordinates is given. Normally Point.
	Type PlaceElemGeometryGeojsonType `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
}

type PlaceElemGeometryGeojsonType string

const PlaceElemGeometryGeojsonTypePoint PlaceElemGeometryGeojsonType = "Point"

type ProductElem struct {
	// The code for the subject in a scheme (= controlled vocabulary) which is
	// identified by the scheme property. "FTFRI", "TTNJE"
	Code string `json:"code,omitempty" yaml:"code,omitempty" mapstructure:"code,omitempty"`

	// The name of the product code. "Feature Fritid", "Nyheter Nöje", etc
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// The identifier of a scheme (= controlled vocabulary) which includes a code for
	// the product. http://tt.se/spec/product/1.0/
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty" mapstructure:"scheme,omitempty"`
}

type Profile string

const (
	ProfileDATA Profile = "DATA"
	ProfileINFO Profile = "INFO"
	ProfilePUBL Profile = "PUBL"
	ProfileRAW  Profile = "RAW"
)

type Pubstatus string

const (
	PubstatusCanceled     Pubstatus = "canceled"
	PubstatusCommissioned Pubstatus = "commissioned"
	PubstatusReplaced     Pubstatus = "replaced"
	PubstatusUsable       Pubstatus = "usable"
	PubstatusWithheld     Pubstatus = "withheld"
)

// Wrapper for different renditions of non-textual content of the news object
type Renditions map[string]Rendition

// Rendition identifier.
type Rendition struct {
	// Href is the URL for accessing the rendition as a resource.
	Href string `json:"href,omitempty"`
	// Mimetype which applies to the rendition.
	Mimetype string `json:"mimetype,omitempty"`
	// Title for the link to the rendition resource.
	Title string `json:"title,omitempty"`
	// Height - for still and moving images: the height of the display area
	// measured in $$TT: unit and defaults to pixels.
	Height int `json:"height,omitempty"`
	// Width - for still and moving images: the width of the display area
	// measured in $$TT: unit and defaults to pixels.
	Width int `json:"width,omitempty"`
	// SizeInBytes of the the rendition resource.
	SizeInBytes int `json:"sizeinbytes,omitempty"`
	// Usage - $$TT: One of 'Thumbnail', 'Preview', 'Hires' or 'Hidef'.
	Usage string `json:"usage,omitempty"`
	// Variant - $$TT: One of 'Normal', 'Watermark', 'BlackAndWhite',
	// 'Cropped' or 'Framegrab'.
	Variant string `json:"variant,omitempty"`
	// Unit - $$TT: The unit for width/height. Either px or mm.
	Unit string `json:"unit,omitempty"`
	// Bitrate - $$TT: Video bitrate (if video).
	Bitrate string `json:"bitrate,omitempty"`
	// Duration of the content in seconds. (Added in version 1.2. Issue #18). nar:remoteContent@duration  $$TT: Video clip curation in seconds.
	Duration float64 `json:"duration,omitempty"`
	// Format - binary format name. (Added in version 1.2. Issue #18). nar:remoteContent@format.
	Format string `json:"format,omitempty"`
	// PrintSize - calculated size of a 300 dpi upsampled image.
	PrintSize float64 `json:"printsize,omitempty"`
}

type Representationtype string

const (
	RepresentationtypeAssociated Representationtype = "associated"
	RepresentationtypeComplete   Representationtype = "complete"
	RepresentationtypeIncomplete Representationtype = "incomplete"
)

type RevisionsElem struct {
	// $$TT: Array of identifiers this revision is replacing.
	Replacing []string `json:"replacing,omitempty" yaml:"replacing,omitempty" mapstructure:"replacing,omitempty"`

	// $$TT: Short name given to article while in production.
	Slug string `json:"slug,omitempty" yaml:"slug,omitempty" mapstructure:"slug,omitempty"`

	// $$TT: The identifier of the previous revision.
	Uri string `json:"uri" yaml:"uri" mapstructure:"uri"`

	// Date and time when this version was published = created. (Added in 1.4)
	Versioncreated *time.Time `json:"versioncreated,omitempty" yaml:"versioncreated,omitempty" mapstructure:"versioncreated,omitempty"`
}

// Expression of rights to be applied to content. nar:rightsInfo (Added in 1.4)
type Rightsinfo struct {
	// Contains a rights expression as defined by a Rights Expression Language.
	// nar:rightsExpressionXML or nar:rightsExpressionData
	Encodedrights string `json:"encodedrights,omitempty" yaml:"encodedrights,omitempty" mapstructure:"encodedrights,omitempty"`

	// Identifier for the Rights Expression language used. nar:@langid
	Langid string `json:"langid,omitempty" yaml:"langid,omitempty" mapstructure:"langid,omitempty"`

	// A link from the current Item to Web resource with rights related information.
	// nar:link
	Linkedrights string `json:"linkedrights,omitempty" yaml:"linkedrights,omitempty" mapstructure:"linkedrights,omitempty"`
}

type Sector string

const (
	SectorEKO Sector = "EKO"
	SectorFEA Sector = "FEA"
	SectorINR Sector = "INR"
	SectorKLT Sector = "KLT"
	SectorNOJ Sector = "NOJ"
	SectorPRM Sector = "PRM"
	SectorSPT Sector = "SPT"
	SectorUTR Sector = "UTR"
)

// $$TT: signals is suggested by AP but not yet included in ninjs. When included it
// will probably hold a large number of properties.
type Signals struct {
	// $$TT: Array of tags set for this delivery
	Deliverytags []string `json:"deliverytags,omitempty" yaml:"deliverytags,omitempty" mapstructure:"deliverytags,omitempty"`

	// $$TT: Number of pages in this delivery.
	Multipagecount *float64 `json:"multipagecount,omitempty" yaml:"multipagecount,omitempty" mapstructure:"multipagecount,omitempty"`

	// $$TT: Code for this page product
	Pagecode string `json:"pagecode,omitempty" yaml:"pagecode,omitempty" mapstructure:"pagecode,omitempty"`

	// $$TT: What type of page product. An abbreviation like IURDAG.
	Pageproduct string `json:"pageproduct,omitempty" yaml:"pageproduct,omitempty" mapstructure:"pageproduct,omitempty"`

	// $$TT: Variant of this page product
	Pagevariant string `json:"pagevariant,omitempty" yaml:"pagevariant,omitempty" mapstructure:"pagevariant,omitempty"`

	// $$TT: Array of pagenumbers for the pages in this delivery. (A pagenumber can
	// also be a letter.)
	Paginae []string `json:"paginae,omitempty" yaml:"paginae,omitempty" mapstructure:"paginae,omitempty"`

	// $$TT: If true this is a retransmission without content change. Also called OMS
	// in the slugline. If the signal do not exist in an item it means retransmission
	// is false.
	Retransmission *bool `json:"retransmission,omitempty" yaml:"retransmission,omitempty" mapstructure:"retransmission,omitempty"`

	// $$TT: If this item is an update of an earlier item this signal indicate what
	// type of update it is. UV mean that the story have developed or changed. KORR
	// mean that some spelling or grammar have been corrected. RA that some fact have
	// been corrected. The connection to earlier item(s) is found in replacing and
	// revisions.
	Updatetype *DocumentSignalsUpdatetype `json:"updatetype,omitempty" yaml:"updatetype,omitempty" mapstructure:"updatetype,omitempty"`
}

type DocumentSignalsUpdatetype string

const (
	DocumentSignalsUpdatetypeKORR DocumentSignalsUpdatetype = "KORR"
	DocumentSignalsUpdatetypeRA   DocumentSignalsUpdatetype = "RÄ"
	DocumentSignalsUpdatetypeUV   DocumentSignalsUpdatetype = "UV"
)

// An object with information about standard, version and schema this instance is
// valid against. nar:standard, nar:standardversion and xml:schema issue #43.
// (Added in version 1.3)
type Standard struct {
	// For example ninjs. nar:standard
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// The uri of the json schema to use for validation.
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty" mapstructure:"schema,omitempty"`

	// For example 1.3. nar:standardversion
	Version string `json:"version,omitempty" yaml:"version,omitempty" mapstructure:"version,omitempty"`
}

type SubjectElem struct {
	// The code for the subject in a scheme (= controlled vocabulary) which is
	// identified by the scheme property
	Code string `json:"code,omitempty" yaml:"code,omitempty" mapstructure:"code,omitempty"`

	// The confidence with which the metadata has been assigned.
	Confidence *int `json:"confidence,omitempty" yaml:"confidence,omitempty" mapstructure:"confidence,omitempty"`

	// Specifies which entity (person, organisation or system) that has created or
	// last edited the property.
	Creator string `json:"creator,omitempty" yaml:"creator,omitempty" mapstructure:"creator,omitempty"`

	// The name of the subject
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// The relationship of the content of the news object to the subject
	Rel string `json:"rel,omitempty" yaml:"rel,omitempty" mapstructure:"rel,omitempty"`

	// The relevance of the metadata to the news content to which it is attached.
	Relevance *int `json:"relevance,omitempty" yaml:"relevance,omitempty" mapstructure:"relevance,omitempty"`

	// The identifier of a scheme (= controlled vocabulary) which includes a code for
	// the subject. $$TT: http://tt.se/spec/subref/1.0/ http://tt.se/spec/keyword/1.0/
	// http://tt.se/spec/eventtype/1.0/
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty" mapstructure:"scheme,omitempty"`
}

type TrustindicatorElem struct {
	// The code for the trust indicator in a scheme (= controlled vocabulary) which is
	// identified by the scheme property
	Code string `json:"code,omitempty" yaml:"code,omitempty" mapstructure:"code,omitempty"`

	// The URL for accessing the trust indicator resource.
	Href string `json:"href,omitempty" yaml:"href,omitempty" mapstructure:"href,omitempty"`

	// The identifier of a scheme (= controlled vocabulary) which includes a code for
	// the trust indicator
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty" mapstructure:"scheme,omitempty"`

	// The title of the resource being referenced.
	Title string `json:"title,omitempty" yaml:"title,omitempty" mapstructure:"title,omitempty"`
}

type Type string

const (
	TypeAudio     Type = "audio"
	TypeComponent Type = "component"
	TypeComposite Type = "composite"
	TypeEvent     Type = "event"
	TypeGraphic   Type = "graphic"
	TypePicture   Type = "picture"
	TypePlanning  Type = "planning"
	TypeText      Type = "text"
	TypeVideo     Type = "video"
)

var enumValues_AdviceElemRole = []interface{}{
	"publish",
}

var enumValues_PlaceElemGeometryGeojsonType = []interface{}{
	"Point",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PlaceElemGeometryGeojsonType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PlaceElemGeometryGeojsonType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PlaceElemGeometryGeojsonType, v)
	}
	*j = PlaceElemGeometryGeojsonType(v)
	return nil
}

var enumValues_Type = []interface{}{
	"text",
	"audio",
	"video",
	"picture",
	"graphic",
	"composite",
	"planning",
	"component",
	"event",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Type) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_Type {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_Type, v)
	}
	*j = Type(v)
	return nil
}

var enumValues_Profile = []interface{}{
	"PUBL",
	"DATA",
	"INFO",
	"RAW",
}

var enumValues_Pubstatus = []interface{}{
	"usable",
	"withheld",
	"canceled",
	"replaced",
	"commissioned",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Pubstatus) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_Pubstatus {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_Pubstatus, v)
	}
	*j = Pubstatus(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Profile) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_Profile {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_Profile, v)
	}
	*j = Profile(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Sector) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_Sector {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_Sector, v)
	}
	*j = Sector(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AdviceElemRole) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AdviceElemRole {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AdviceElemRole, v)
	}
	*j = AdviceElemRole(v)
	return nil
}

var enumValues_Representationtype = []interface{}{
	"complete",
	"incomplete",
	"associated",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Representationtype) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_Representationtype {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_Representationtype, v)
	}
	*j = Representationtype(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RevisionsElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["uri"]; !ok || v == nil {
		return fmt.Errorf("field uri in RevisionsElem: required")
	}
	type Plain RevisionsElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RevisionsElem(plain)
	return nil
}

var enumValues_Sector = []interface{}{
	"INR",
	"UTR",
	"EKO",
	"KLT",
	"SPT",
	"FEA",
	"NOJ",
	"PRM",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Document) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["uri"]; !ok || v == nil {
		return fmt.Errorf("field uri in Document: required")
	}
	type Plain Document
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Document(plain)
	return nil
}
