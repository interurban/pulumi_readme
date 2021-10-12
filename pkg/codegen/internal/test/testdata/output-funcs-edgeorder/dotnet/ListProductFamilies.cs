// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Myedgeorder
{
    public static class ListProductFamilies
    {
        /// <summary>
        /// The list of product families.
        /// API Version: 2020-12-01-preview.
        /// </summary>
        public static Task<ListProductFamiliesResult> InvokeAsync(ListProductFamiliesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListProductFamiliesResult>("myedgeorder::listProductFamilies", args ?? new ListProductFamiliesArgs(), options.WithVersion());
    }


    public sealed class ListProductFamiliesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details
        /// </summary>
        [Input("customerSubscriptionDetails")]
        public Inputs.CustomerSubscriptionDetails? CustomerSubscriptionDetails { get; set; }

        /// <summary>
        /// $expand is supported on configurations parameter for product, which provides details on the configurations for the product.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        [Input("filterableProperties", required: true)]
        private Dictionary<string, ImmutableArray<Inputs.FilterableProperty>>? _filterableProperties;

        /// <summary>
        /// Dictionary of filterable properties on product family.
        /// </summary>
        public Dictionary<string, ImmutableArray<Inputs.FilterableProperty>> FilterableProperties
        {
            get => _filterableProperties ?? (_filterableProperties = new Dictionary<string, ImmutableArray<Inputs.FilterableProperty>>());
            set => _filterableProperties = value;
        }

        /// <summary>
        /// $skipToken is supported on list of product families, which provides the next page in the list of product families.
        /// </summary>
        [Input("skipToken")]
        public string? SkipToken { get; set; }

        public ListProductFamiliesArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListProductFamiliesResult
    {
        /// <summary>
        /// Link for the next set of product families.
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// List of product families.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProductFamilyResponse> Value;

        [OutputConstructor]
        private ListProductFamiliesResult(
            string? nextLink,

            ImmutableArray<Outputs.ProductFamilyResponse> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}